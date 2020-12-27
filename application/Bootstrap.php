<?php
/**
 * 框架启动加载配置文件
 *
 */
class Bootstrap extends Yaf_Bootstrap_Abstract{

    /**
     * 注册本地类名前缀, 这部分类名将会在本地类库查找
     *
     * @param Yaf_Dispatcher $dispatcher
     */
	public function _initLoader(Yaf_Dispatcher $dispatcher){
		Yaf_Loader::getInstance()->registerLocalNameSpace(
			array('Helper'));
    }

    /**
     * 初始化配置信息
     *
     * @param Yaf_Dispatcher $dispatcher
     */
	public function _initConfig(Yaf_Dispatcher $dispatcher) {
        Yaf_Registry::set("config", Yaf_Application::app()->getConfig());
		Comm_Db::auto_configure_pool();
	}

	/**
	 * 关闭自动渲染, 因为我们都是自己主动调用视图的render
	 *
	 * @param Yaf_Dispatcher $dispatcher
	 */
	public function _initView(Yaf_Dispatcher $dispatcher){
		$dispatcher->disableView();
	}

	/**
	 * 初始化插件
	 *
	 * @param Yaf_Dispatcher $dispatcher
	 */
	public function _initPlugin(Yaf_Dispatcher $dispatcher) {
		if (defined('MODULE_NAME')) {
		    $className = MODULE_NAME . CONFIG_SUFFIX_PLUGIN;
		    $rewrite = new $className();
		    $dispatcher->registerPlugin($rewrite);
		}
    }
}