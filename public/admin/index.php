<?php
/**
 * api入口文件, 供内部访问
 *
 */
require_once dirname(dirname(__FILE__)) . '/init.php';
define('MODULE_NAME', 'Admin');
$app = new Yaf_Application(CONFIG_PATH_CONFIG . '/application.ini');
$app->bootstrap()->run();