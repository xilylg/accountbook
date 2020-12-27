<?php
/**
 *
 * @param unknown_type $code
 * @param unknown_type $msg
 * @param unknown_type $data
 **/
class Helper_Render {
    public static function renderAjax($code, $msg = '', $data = NULL) {
        if (!headers_sent()) {
            header('Content-type: application/json; charset=utf-8', TRUE);
        }
        if ($data === NULL) {
        	$data = new stdClass();
        }
        self::outJson($code, $msg, $data);
    }
    
    /**
     * 按json格式输出响应
     *
     * @param string|int	$code			js的错误代码/行为代码
     * @param string		$message		可选。行为所需文案或者错误详细描述。默认为空。
     * @param mixed			$data			可选。附加数据。
     * @return void
     */
    public static function outJson($code, $message = '', $data = NULL) {
    		$output_data = array(
    				"code" => $code,
    				"msg" => strval($message),
    				"data" => $data,
    		);
    		echo json_encode(array_merge(self::$meta,$output_data));
    }
}

