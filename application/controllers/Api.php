<?php
class ApiController extends AbstractController
{
	//0无须检查登录，1必须登录
	protected  $login = 1;
	
	public function init()
	{
		
	}
	
	private function login()
	{
		
		if ($this->login == 1) {
			
		}
	}
	
	private function getCurrentUserinfo()
	{
		return [
				'uid' => 1,
				'nickname' => 'Sandy',
				'gender' => 1,
		];
	}
}
