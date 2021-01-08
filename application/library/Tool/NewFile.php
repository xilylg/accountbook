<?php

$tpl = '<?php

class {%s}Model extends AbstractModel
{
	public function __construct($log = NULL) {
	    parent::__construct($log);
		$this->datamodel = new {%s}Data($this->log);
	}
}';

$files = ['flow', 'item', 'member', 'target'];
$path = "/Users/xiaoyu45/accountbook/application/models/";
foreach($files as $file) {
    $cur_tpl = str_replace("{%s}", $file, $tpl);
    file_put_contents($path.$file.".php", $cur_tpl);
}

