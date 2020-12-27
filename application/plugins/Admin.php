<?php

class AdminPlugin extends Yaf_Plugin_Abstract
{
    public function routerShutdown(Yaf_Request_Abstract $request, Yaf_Response_Abstract $response)
    {
        $request->setModuleName(MODULE_NAME);
    }

}