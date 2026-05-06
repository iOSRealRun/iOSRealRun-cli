"""
Shared runtime configuration for command helpers.
"""

import os
import sys

from tools.config import config


def get_os():
    platform = sys.platform
    if -1 != platform.find("win32"):
        return "win"
    if -1 != platform.find("darwin"):
        return "darwin"
    return "linux"


OS = get_os()
separator = {"win": "\\", "darwin": "/", "linux": "/"}[OS]
seperator = separator  # Backward-compatible spelling used elsewhere.
libimobiledevice_dir = config.libimobiledeviceDir + separator + OS
libimobiledeviceDir = libimobiledevice_dir
v = config.v


def _library_env(variable_name):
    result = os.environ.copy()
    result[variable_name] = os.getcwd() + "/" + libimobiledevice_dir
    return result


env = {
    "win": None,
    "darwin": _library_env("DYLD_LIBRARY_PATH"),
    "linux": _library_env("LD_LIBRARY_PATH"),
}
