<?php

/**
 * 后台公共文件 
 * @file   Common.php  
 * @date   2016-8-24 18:28:34 
 * @author Zhenxun Du<5552123@qq.com>  
 * @version    SVN:$Id:$ 
 */

namespace app\admin\controller;

use think\Controller;
use think\Session;

class Common extends Controller {

    public function _initialize() {
        
        $username=Session::get('username');
        //var_dump($lists);
        if (!$username) $this->error("还未登录");
        $this->assign('username',$username);
    }

}