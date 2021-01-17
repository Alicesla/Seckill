<?php
namespace app\admin\controller;

use think\Controller;
use think\Loader;
use think\Db;
use think\Session;
class Main extends Common {
    
    public function index() {
        return $this->fetch();
    }
    public function qianggou() {
        return $this->fetch();
    }
    public function res() {
        
        return $this->fetch();
    }
}?>
