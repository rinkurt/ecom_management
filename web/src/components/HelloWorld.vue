<template>
  <div class="hello">
    <el-container style="height: 100%; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
        <el-menu :default-openeds="['1', '3']">
          <el-submenu index="1">
            <template slot="title"
            ><i class="el-icon-message"></i>运营管理
            </template
            >
            <el-menu-item-group>
<!--              <template slot="title">分组一</template>-->
              <el-menu-item index="1-1">视频管理</el-menu-item>
              <el-menu-item index="1-2">用户管理</el-menu-item>
              <el-menu-item index="1-2">商品管理</el-menu-item>
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
          <el-form :inline="true" ref="form" :model="form" label-width="80px" :label-position="'right'"
                   style="text-align: left">
            <el-form-item label="视频ID" prop="item_id">
              <el-input v-model.number="form.item_id" type="number"></el-input>
            </el-form-item>
            <el-form-item label="标题" prop="title">
              <el-input v-model="form.title"></el-input>
            </el-form-item>
            <el-form-item label="创建时间">
              <el-col :span="11">
                <el-date-picker type="date" placeholder="开始" v-model="form.create_time_from"
                                style="width: 100%;"></el-date-picker>
              </el-col>
              <el-col :span="11">
                <el-date-picker type="date" placeholder="结束" v-model="form.create_time_to"
                                style="width: 100%;"></el-date-picker>
              </el-col>
            </el-form-item>
            <el-form-item label="标签" prop="title">
              <el-input v-model="form.label"></el-input>
            </el-form-item>
            <el-form-item label="用户ID" prop="title">
              <el-input v-model.number="form.user_id" type="number"></el-input>
            </el-form-item>
            <el-form-item label="status" prop="title">
              <el-input v-model.number="form.status" type="number"></el-input>
            </el-form-item>
            <el-form-item label="rate" prop="title">
              <el-input v-model.number="form.rate" type="number"></el-input>
            </el-form-item>
            <el-form-item label="内容等级" prop="title">
              <el-input v-model.number="form.content_level" type="number"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmit">搜索</el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-form>
          <el-table :data="tableData" @sort-change='sortChange'>
            <el-table-column sortable='custom' prop="item_id" label="视频ID" width="140">
            </el-table-column>
            <el-table-column sortable='custom' prop="user_id" label="用户ID"></el-table-column>
            <el-table-column prop="title" label="标题" width="120">
            </el-table-column>
            <el-table-column prop="video_url" label="video_url"></el-table-column>
            <el-table-column prop="label" label="标签"></el-table-column>
            <el-table-column prop="status" label="状态"></el-table-column>
            <el-table-column prop="rate" label="推荐状态"></el-table-column>
            <el-table-column prop="content_level" label="内容等级"></el-table-column>
            <el-table-column prop="is_ecom" label="是否电商" :formatter="formatBool"></el-table-column>
            <el-table-column sortable='custom' prop="create_time" label="创建时间" :formatter="formatDate"></el-table-column>
          </el-table>
          <el-pagination
            background
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-size="size"
            layout="total, prev, pager, next, jumper"
            :total="total"
            style="margin-top: 20px">
          </el-pagination>
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
import request from "@/utils/request";

export default {
  data() {
    return {
      tableData: [],
      total: 0,
      currentPage: 1,
      size: 10,
      orderBy: "create_time",
      order: 0,
      form: {
        item_id: null,
        title: null,
        create_time_from: null,
        create_time_to: null,
        label: null,
        user_id: null,
        status: null,
        rate: null,
        content_level: null
      }
    }
  },

  mounted: function () {
    window.vue = this
    this.getItemList()
  },

  methods: {
    onSubmit() {
      this.currentPage = 1
      this.getItemList()
    },

    resetForm() {
      this.$refs['form'].resetFields();
    },

    handleCurrentChange(currentPage) {
      this.currentPage = currentPage
      this.getItemList()
    },

    sortChange(column) {
      this.orderBy = column.prop
      if (column.order === 'descending') {
        this.order = 0
      } else {
        this.order = 1
      }
      this.getItemList()
    },

    getItemList() {
      let ctx = this
      let req = {
        offset: (this.currentPage - 1) * this.size,
        size: this.size,
        order_by: this.orderBy,
        order: this.order,
      }

      if (this.form.item_id != null) {
        req.item_id = [this.form.item_id]
      }
      if (this.form.title != null) {
        req.title = this.form.title
      }
      if (this.form.create_time_from != null) {
        req.create_time_from = this.form.create_time_from.getTime()
      }
      if (this.form.create_time_to != null) {
        req.create_time_to = this.form.create_time_to.getTime() + 86399
      }
      if (this.form.label != null) {
        req.label = [this.form.label]
      }
      if (this.form.user_id != null) {
        req.user_id = [this.form.user_id]
      }
      if (this.form.status != null) {
        req.status = [this.form.status]
      }
      if (this.form.rate != null) {
        req.rate = [this.form.rate]
      }
      if (this.form.content_level != null) {
        req.content_level = [this.form.content_level]
      }

      request.post('/item/get_item_list', req)
        .then((res) => {
          if (res.data.code === 0) {
            ctx.tableData = res.data.data.item_list
            ctx.total = res.data.data.total
          } else {
            ctx.popWarning(res.data.msg)
          }
        })
        .catch((err) => {
          ctx.popWarning(err)
        })
    }
  }
}
</script>
