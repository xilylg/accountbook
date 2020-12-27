<?php

class ApiPlugin extends Yaf_Plugin_Abstract
{

    public function routerShutdown(Yaf_Request_Abstract $request, Yaf_Response_Abstract $response)
    {
        $request->setModuleName(MODULE_NAME);
        $controllerName = Yaf_Dispatcher::getInstance()->getRequest()->getControllerName();
        
        //cardlist跳转, 232118_-_cardlist_-_{$controller}
        if ($controllerName == 'Pub_Recom' && isset($_REQUEST['containerid']) && strpos($_REQUEST['containerid'], '232118_-_cardlist_-_') === 0) {
            $containerids = explode("_-_", $_REQUEST['containerid']);
            if (!empty($containerids[2])) {
                $controllerName = 'Cardlist_' . ucfirst($containerids[2]);
                $request->setControllerName($controllerName);
            }
        }
    }
}