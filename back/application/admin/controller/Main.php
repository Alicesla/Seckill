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
    public function tianjiashangpin() {
        $role=Session::get('userrole');
        $info_1 = db('roleaccess')->where('access', 'tianjiashangpin')->find();
        if ($info_1[$role]==0){
           $this->tiaochu();
        }
        return $this->fetch();
    }
    public function shangpinxinxi(){
        $lists = db('boost')->select();
        $this->assign('lists',$lists);
        return $this->fetch();
    }

    public function shanchushangpin() {
        $role=Session::get('userrole');
        $info_1 = db('roleaccess')->where('access', 'shanchushangpin')->find();
        if ($info_1[$role]==0){
          $this->tiaochu();
        }
        return $this->fetch();
    }
    
    public function shangpinshuliang() {
        $role=Session::get('userrole');
        $info_1 = db('roleaccess')->where('access', 'shangpinshuliang')->find();
        if ($info_1[$role]==0){
           $this->tiaochu();
        }
        return $this->fetch();
    }
    public function shangpinshijian() {
        $role=Session::get('userrole');
        $info_1 = db('roleaccess')->where('access', 'shangpinshijian')->find();
        if ($info_1[$role]==0){
           $this->tiaochu();
        }
        return $this->fetch();
    }
    public function shangpinjiage() {
        $role=Session::get('userrole');
        $info_1 = db('roleaccess')->where('access', 'shangpinjiage')->find();
        if ($info_1[$role]==0){
           $this->tiaochu();
        }
        return $this->fetch();
    }
    public function qianggoushumu() {
        $role=Session::get('userrole');
        $info_1 = db('roleaccess')->where('access', 'qianggoushumu')->find();
        if ($info_1[$role]==0){
           $this->tiaochu();
        }
        return $this->fetch();
    }
    public function gerenxinxi() {
        $username=Session::get('username');
        $sql="select * from userinfo where username='$username'";
        $lists=Db::query($sql);
        $this->assign('lists',$lists);
        return $this->fetch();
    }
    public function xiugaimima() {
        return $this->fetch();
    }
    
    
    
     public function tianjia_shangpin() {
        $name=input('post.name');
        $num=input('post.num');
        $price=input('post.price');
        $time=input('post.time');
        $maxnums=input('post.maxnums');
        $info = db('boost')->where('name', $name)->find();
        if (!$info){
            $data = [
                'name' => $name,
                'num' => $num,
                'price'=>$price,
                'time'=>$time,
                'max_nums'=>$maxnums,
            ];
            $res = Db::name('boost')->insert($data);
        }
        else {
            $this->error('该商品已存在');
        }
        return $this->success('添加成功','main/shangpinxinxi');
        
    }

    public function shanchu_shangpin() {
        $name=input('post.name');
        
        $info = db('boost')->where('name', $name)->find();
        if (!$info){
            $this->error('该商品不存在');
        }
        else {
            $res = Db::name('boost')->where('name',$name)->delete();
            return $this->success('删除成功','main/shangpinxinxi');   
        }
    }
    public function shangpin_shuliang(){
        $name = input('post.name');
        $num=input('post.num');
        $info = db('boost')->field('name','num')->where('name', $name)->find();

        
        if (!$info) {
            $this->error('该商品不存在','main/shangpinxinxi');
        }
        $res = Db::name('boost')->where('name',$name)->update(['num'=>$num]);
        
        $this->success('修改商品数量成功', 'main/shangpinxinxi');
        
    }
    public function shangpin_shijian(){
        $name = input('post.name');
        $time=input('post.time');
        $info = db('boost')->field('name')->where('name', $name)->find();

        
        if (!$info) {
            $this->error('该商品不存在','main/shangpinxinxi');
        }
        $res = Db::name('boost')->where('name',$name)->update(['time'=>$time]);
        
        $this->success('修改抢购时间成功', 'main/shangpinxinxi');
        
    }
    public function shangpin_jiage(){
        $name = input('post.name');
        $price=input('post.price');
        $info = db('boost')->field('name')->where('name', $name)->find();

        
        if (!$info) {
            $this->error('该商品不存在','main/shangpinxinxi');
        }
        $res = Db::name('boost')->where('name',$name)->update(['price'=>$price]);
        
        $this->success('修改商品价格成功', 'main/shangpinxinxi');
        
    }
    public function qianggou_shumu(){
        $name = input('post.name');
        $maxnums=input('post.maxnums');
        $info = db('boost')->field('name')->where('name', $name)->find();
        
        if (!$info) {
            $this->error('该商品不存在','main/shangpinxinxi');
        }
        $res = Db::name('boost')->where('name',$name)->update(['max_nums'=>$maxnums]);
        
        $this->success('修改最大抢购数目成功', 'main/shangpinxinxi');
        
    }
    public function xiugai_mima() {
        $md5_salt = config('md5_salt');
        $username=input('post.username');
        $password_old=input('post.password_old');
        $password_new=input('post.password_new');
        $info = db('userinfo')->field('username','password')->where('username', $username)->find();
        if ($username!=Session::get('username')){
            $this->error('您无权限修改其他用户密码');
        }
        if (!$info) {
            $this->error('用户名或密码错误');
        }

        if (md5(md5($password_old).$md5_salt) != $info['password']) {
            $this->error('用户名或密码错误');
        } 
        $res = Db::name('userinfo')->where('username',$username)->update(['password'=>md5(md5($password_new).$md5_salt)]);
        $this->success('修改密码成功', 'main/index');
    }
    private function tiaochu(){
         ?> 
            <script type="text/javascript">
                alert("您无权限进行该操作！");
                window.location.href="./index";
            </script>
        <?php 
    }
   
    
}?>
