<template>
  <div class="recommend">
    <el-container style="height: 100%; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
        <el-menu :default-openeds="['1', '3']">
          <el-submenu index="1">
<!--            <template slot="title"-->
<!--            ><i class="el-icon-message"></i>===-->
<!--            </template-->
<!--            >-->
            <el-menu-item-group>
<!--              <template slot="title">分组一</template>-->
              <el-menu-item index="1-1">视频推荐</el-menu-item>
              <el-menu-item index="1-2">商品浏览</el-menu-item>
<!--              <el-menu-item index="1-2">商品管理</el-menu-item>-->
            </el-menu-item-group>
<!--            <el-menu-item-group title="分组2">-->
<!--              <el-menu-item index="1-3">选项3</el-menu-item>-->
<!--            </el-menu-item-group>-->
<!--            <el-submenu index="1-4">-->
<!--              <template slot="title">选项4</template>-->
<!--              <el-menu-item index="1-4-1">选项4-1</el-menu-item>-->
<!--            </el-submenu>-->
          </el-submenu>
        </el-menu>
      </el-aside>

      <el-container>
        <el-header style="text-align: right; font-size: 12px">
          <el-dropdown>
            <i class="el-icon-setting" style="margin-right: 15px"></i>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>查看</el-dropdown-item>
              <el-dropdown-item>新增</el-dropdown-item>
              <el-dropdown-item>删除</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
          <span>用户</span>
        </el-header>

        <el-main>
          <div class="block">
            <el-image style="width: 300px" src="https://s1.ax1x.com/2022/08/04/veMOSg.jpg"></el-image>
          </div>
          <div>
            标题：这款杯子强烈推荐
          </div>
          <div>
            发布者：user1
          </div>
          <el-row>
            <el-button>点赞：1</el-button>
            <el-button>查看商品</el-button>
            <el-button>查看发布者</el-button>
            <el-button>下一条</el-button>
          </el-row>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style>
.el-header {
  background-color: #b3c0d1;
  color: #333;
  line-height: 60px;
}

.el-aside {
  color: #333;
}
</style>

<script>
export default {
  name: 'Recommend',
}
</script>
