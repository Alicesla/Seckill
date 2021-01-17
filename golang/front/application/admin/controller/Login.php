<?php

/**
 *  登陆页
 * @file   Login.php  
 * @date   2016-8-23 19:52:46 
 * @author Zhenxun Du<5552123@qq.com>  
 * @version    SVN:$Id:$ 
 */

namespace app\admin\controller;

use think\Controller;
use think\Loader;
use think\Session;
class Login extends Controller {

    /**
     * 登入
     */
    public function index() {
            
            return $this->fetch();
        
    }
    
    /**
     * 处理登录
     */
    public function dologin() {
        //验证密码流程
        
        //假设用户名密码争取
        $username = input('post.username');
        $password = input('post.password');
        $info = db('userinfo')->field('username,password,userrole')->where('username', $username)->find();
        
        if (!$info) {
            $this->error('用户名或密码错误');
        }
        $md5_salt = config('md5_salt');
        if (md5(md5($password).$md5_salt) != $info['password']) {
            $this->error('用户名或密码错误');
        } 
        else {
            Session::set('username', $info['username']);
            Session::set('userrole', $info['userrole']);
             //记录登录信息
            $this->success('登入成功', 'main/index');
        }
    }

    /**
     * 登出
     */
    public function logout() {
        Session::set('username', null);
        $this->success('退出成功', 'login/index');
    }

    public function test(){
        echo 'test';
    }

}
