<?php
class AbstractController extends Yaf_Controller_Abstract
{
	protected $method = 'request';
	protected $verify = [];
	protected $filter = [];
	
	public function init()
	{
		$this->getCommonParams();
		
		if　(!empty($this->verify)) {
			$this->verifyParams();
		}
	}
	
	public function getCommonParams()
	{
	}
	
	public function verifyParams()
    {
    	if (empty($this->verify)) {
            return true;
        }
        
        foreach ($this->verify as $name => $rules) {
        	$value = 
            $verify_result = Context_Verify_Verify::request($name, $value, $rules);
            if ($verify_result['code'] == 0) {
            	return $this->error('param_error', $verify_result['error_msg']);
            }
            $this->params[$name] = $value;
        }
        return true;
    }
    
    public function success($data)
    {
    	$code  = Helper_Conf::get("interfacecode.common.success");
    	$msg = Helper_Conf::get("interfacecode.common.success_text");
        return Helper_Render::renderAjax(Helper_Conf::get("interfacecode.common.success"), "ok", $data);
    }
    
    public function error($err_code = "error", $msg = "")
    {
    	if (empty($msg)) {
    		$msg =  Helper_Conf::get("interfacecode.common.{$err_code}_text");
    	}
        $err_code = empty($err_code) ? Helper_Conf::get("interfacecode.common.sys_error") : $err_code;
        Helper_Render::renderAjax($err_code, $msg);
    }
}