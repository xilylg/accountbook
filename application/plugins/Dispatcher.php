<?php

/**
 * 分发器
 */
class Dispatcher
{

    private $instance = null;

    private $module = null;

    private $controller = null;

    private $action = null;

    public function setModule($module)
    {
        if (! empty($module)) {
            self::$module = $module;
        }
        return self::getInstance();
    }

    public function getModule()
    {
        $module = Yaf_Dispatcher::getInstance()->getRequest()->getModuleName();
        return strtolower($module);
    }

    public function setController($controller = null)
    {
        if (! is_null($controller)) {
            self::$controller = $controller;
        }
        return self::getInstance();
    }

    public function getController()
    {
        return Yaf_Dispatcher::getInstance()->getRequest()->getControllerName();
    }

    public function setAction($action = null)
    {
        if (! is_null($action)) {
            self::$action = $action;
        }
        return self::getInstance();
    }

    public function getAction()
    {
        return Yaf_Dispatcher::getInstance()->getRequest()->getActionName();
    }

    public function getControllerFromPath($path = null)
    {
        $controller = CONFIG_DEFAULT_CONTROLLER;
        if (! empty($path)) {
            $controller = str_replace(' ', '', ucwords(str_replace('_', ' ', $path)));
        }
        return $controller;
    }

    public function getActionFromPath($path = null)
    {
        $action = CONFIG_DEFAULT_ACTION;
        if (! empty($path)) {
            $action = str_replace(' ', '', ucwords(str_replace('_', ' ', $path)));
            $action{0} = strtolower($action{0});
        }
        return $action;
    }
}
