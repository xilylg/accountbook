<?php
ini_set("display_errors", 1);
error_reporting(E_ALL & ~ E_NOTICE & ~ E_WARNING & ~ E_STRICT);
define('ROOTPATH', dirname(__DIR__));
define('APPPATH', ROOTPATH . '/application');
define('CONFIG_PATH_CONFIG', APPPATH . "/config");
define('CONFIG_SUFFIX_PLUGIN', 'Plugin');
define('CONFIG_DEFAULT_ACTION', 'index');
define('CONFIG_DEFAULT_CONTROLLER', 'index');
define('EXT', '.php');
define('LOG_PATH', $_SERVER['log_path']);
include APPPATH . "/plugins/Dispatcher.php";

