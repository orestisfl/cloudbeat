"""
This module provides functions for generating and formatting policy templates
and inputs based on provided data and templates.
"""

import copy
import uuid
from typing import Dict, Tuple
from urllib.parse import parse_qs, unquote, urlparse

from fleet_api.common_api import get_package
from loguru import logger
from munch import Munch
from packaging import version

# Constants
CLOUD_SECURITY_POSTURE = "cloud_security_posture"
REQUIRE_VARS = ["cloudbeat/cis_aws", "cloudbeat/cis_eks"]
DEFAULT_PACKAGE_POLICY_TEMPLATE = {
    "policy_id": "",
    "package": {},
    "name": "",
    "description": "",
    "namespace": "default",
    "inputs": {},
    "vars": {},
}
SIMPLIFIED_AGENT_POLICY = {
    "name": "",
    "namespace": "default",
    "monitoring_enabled": ["logs", "metrics"],
}
VERSION_MAP = {
    "asset_inventory_aws": "0.2.1",
    "asset_inventory_azure": "0.2.1",
    "cis_k8s": "1.1.0",
    "cis_eks": "1.2.0",
    "cis_aws": "1.2.0",
    "vuln_mgmt_aws": "1.3.0",
    "cis_gcp": "1.5.0",
    "cis_azure": "1.6.0",
}


def generate_input_id(name: str, input_type: str) -> str:
    """
    Generates a unique input ID based on the provided name and input type.

    This function combines the 'name' and 'input_type' parameters to create a unique
    identifier for an input. The resulting ID is useful for organizing and referencing
    inputs in a structured manner.

    Args:
        name (str): The name or identifier associated with the input.
        input_type (str): The type or category of the input.

    Returns:
        str: A unique input ID generated by combining 'name' and 'input_type'.

    Example:
        If 'name' is "cspm" and 'input_type' is "cloudbeat/cis_gcp", calling
        'generate_input_id(name, input_type)' will return "cspm-cloudbeat/cis_gcp".
    """
    return f"{name}-{input_type}"


def format_inputs(policy_templates: list, stream_prefix: str) -> dict:
    """
    Format inputs based on policy templates.

    Args:
        policy_templates (list): List of policy templates.

    Returns:
        dict: Formatted inputs.
    """
    inputs_dict = {}
    for template in policy_templates:
        name = template.get("name", "")
        data_stream = template.get("data_streams", [])[0]
        for template_input in template.get("inputs", []):
            input_type = template_input.get("type", "")
            input_dict = {
                "enabled": False,
                "streams": {
                    f"{stream_prefix}.{data_stream}": {
                        "enabled": False,
                    },
                },
            }
            # Conditionally add "vars" based on input_type
            if input_type in REQUIRE_VARS:
                input_dict["streams"][f"{stream_prefix}.{data_stream}"]["vars"] = {}
            inputs_dict[generate_input_id(name=name, input_type=input_type)] = input_dict
    return inputs_dict


def format_vars(package_vars: list) -> dict:
    """
    Format vars based on package vars.

    Args:
        package_vars (list): List of package vars.

    Returns:
        dict: Formatted vars.
    """
    vars_dict = {}
    for package_var in package_vars:
        vars_dict[package_var.get("name", "")] = ""
    return vars_dict


def update_policy_input_data(data, input_data):
    """
    Recursively updates a dictionary structure with values from another dictionary.

    This function is particularly designed for updating policy-related data structures.

    Args:
        data (dict or list): The dictionary structure to be updated.
        input_data (dict): The dictionary containing values to update 'data' with.
    """
    if isinstance(data, dict):
        for key, value in data.items():
            if key == "enabled":
                data[key] = True
            elif key == "vars" and isinstance(value, dict):
                data[key] = input_data.get("vars", {})
            elif isinstance(value, (dict, list)):
                update_policy_input_data(value, input_data)
    elif isinstance(data, list):
        for item in data:
            update_policy_input_data(item, input_data)


def generate_policy_template(
    cfg: Munch,
    package_name: str = "",
    policy_template: dict = None,
    stream_prefix: str = "",
) -> dict:
    """
    Generate a policy template based on configuration and a template.

    Args:
        cfg (Munch): Configuration data.
        policy_template (dict, optional): Policy template.If not provided,
                                          a default template will be used.

    Returns:
        dict: Generated policy template.
    """
    if policy_template is None:
        policy_template = DEFAULT_PACKAGE_POLICY_TEMPLATE

    generated_policy = copy.deepcopy(policy_template)
    if package_name:
        package_policy_info = get_package(cfg=cfg, package_name=package_name, prerelease=True)
    else:
        package_policy_info = get_package(cfg=cfg, prerelease=True)
    generated_policy["package"] = {
        "name": package_policy_info.get("name", ""),
        "version": package_policy_info.get("version", ""),
    }
    generated_policy["inputs"] = format_inputs(
        package_policy_info.get("policy_templates", []),
        stream_prefix=stream_prefix,
    )
    generated_policy["vars"] = format_vars(package_vars=package_policy_info.get("vars", []))
    return generated_policy


def generate_package_policy(template: dict, policy_input: dict, stream_name: str) -> dict:
    """
    Generate a package policy based on a template and policy input.

    Args:
        template (dict): The package policy template.
        policy_input (dict): The policy input containing values to update.

    Returns:
        dict: The generated package policy.
    """
    package_policy = copy.deepcopy(template)
    integration_key = policy_input.get("input_name", "")
    for input_name, data in package_policy["inputs"].items():
        if integration_key in input_name:
            update_policy_input_data(data, policy_input)
            if "vars" in policy_input and "vars" not in data["streams"][stream_name]:
                data["streams"][stream_name]["vars"] = policy_input["vars"]
    if "posture" in stream_name:
        package_policy["vars"]["posture"] = policy_input.get("posture", "")
        package_policy["vars"]["deployment"] = policy_input.get("deployment", "")
    package_policy["name"] = policy_input.get("name", "")
    return package_policy


def load_data(
    cfg: Munch,
    agent_input: dict,
    package_input: dict,
    stream_name: str,
    package_name: str = "",
) -> Tuple[Dict, Dict]:
    """
    Load agent and package policies based on input data.

    Args:
        cfg (Munch): Configuration data.
        agent_input (dict): Agent policy input data.
        package_input (dict): Package policy input data.

    Returns:
        Tuple[Dict, Dict]: A tuple containing the loaded agent policy and package policy.
    """
    logger.info("Loading agent and package policies")
    agent_policy = SIMPLIFIED_AGENT_POLICY
    agent_policy["name"] = agent_input.get("name", "")
    agent_policy["supports_agentless"] = bool(agent_input.get("supports_agentless"))

    stream_prefix = stream_name.split(".")[0]
    if package_name:
        package_template = generate_policy_template(
            cfg=cfg,
            package_name=package_name,
            stream_prefix=stream_prefix,
        )
    else:
        package_template = generate_policy_template(
            cfg=cfg,
            stream_prefix=stream_prefix,
        )
    package_policy = generate_package_policy(
        template=package_template,
        policy_input=package_input,
        stream_name=stream_name,
    )

    return agent_policy, package_policy


def version_compatible(current_version, required_version):
    """
    Check if the current version is compatible with the required version.

    Args:
        current_version (str): The current version to be checked.
        required_version (str): The required version for compatibility.

    Returns:
        bool: True if the current version is compatible, False otherwise.
    """
    return version.parse(current_version) >= version.parse(required_version)


def generate_random_name(prefix: str) -> str:
    """
    Generate a random name by combining a given prefix with a random suffix.

    Args:
        prefix (str): The prefix to be combined with the random suffix.

    Returns:
        str: The generated name consisting of the prefix and a random 6-character suffix.
    """
    random_uuid = str(uuid.uuid4())
    # Extract the last 6 characters from the UUID
    random_suffix = random_uuid[-6:]
    generated_name = f"{prefix}-{random_suffix}"

    return generated_name


def patch_vars(var_dict, package_version):
    """
    Conditionally updates a dictionary based on the package version.

    This function checks the provided package version and updates the given
    dictionary 'var_dict' with additional fields based on version requirements.

    Args:
        var_dict (dict): The dictionary to be updated.
        package_version (str): The version of the package to determine updates.

    Returns:
        None: This function modifies 'var_dict' in place.
    """
    if version.parse(package_version) >= version.parse("1.5.0"):
        # Add or update fields in the vars_dict based on the version requirements
        var_dict["aws.account_type"] = "single-account"


def get_package_default_url(cfg: Munch, policy_name: str, policy_type: str, package_name: str = "") -> str:
    """
    Get the package default URL for a specific policy and policy type from the configuration.

    Args:
        cfg (Munch): The configuration containing policy information.
        policy_name (str): The name of the policy.
        policy_type (str): The type of the policy.

    Returns:
        str: The default package URL for the specified policy and type.
             An empty string is returned if not found.
    """
    if package_name:
        package_policy = get_package(cfg=cfg, package_name=package_name)
    else:
        package_policy = get_package(cfg=cfg)
    policy_templates = package_policy.get("policy_templates", [])

    for template in policy_templates:
        if template.get("name", "") == policy_name:
            inputs = template.get("inputs", [])

            for policy_input in inputs:
                if policy_input.get("type", "") == policy_type:
                    vars_list = policy_input.get("vars", [])

                    if vars_list:
                        return vars_list[0].get("default", "")

    return ""


def extract_template_url(url_string: str) -> str:
    """
    Extracts the 'templateURL' parameter from a given URL string.

    Args:
        url_string (str): The URL string from which to extract the 'templateURL' parameter.

    Returns:
        str: The value of the 'templateURL' parameter if found in the URL,
             or empty string if the parameter is not present.

    """
    parsed_url = urlparse(url_string, allow_fragments=False)
    query_parameters = parse_qs(parsed_url.query)

    template_url = query_parameters.get("templateURL")

    if not template_url:
        logger.warning("templateURL field is not found")
        return ""

    return template_url[0]


def extract_arm_template_url(url_string: str) -> str:
    """
    Extracts the arm template link from a given URL string.

    Args:
        url_string (str): The URL string from which to extract the arm template link.

    Returns:
        str: The arm template link.

    """
    parsed_url = urlparse(url_string, allow_fragments=False)
    return unquote(parsed_url.path.split("/")[-1]).replace("ACCOUNT_TYPE", "single-account")
