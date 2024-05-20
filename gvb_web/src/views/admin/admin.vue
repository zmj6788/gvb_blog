<template>
  <div class="gvb_admin">
    <GVBAside></GVBAside>

    <div class="main">
      <header>
        <div class="left">
          <a-breadcrumb separator="  ">
            <a-breadcrumb-item>首页</a-breadcrumb-item>
            <a-breadcrumb-item><a href="">个人中心</a></a-breadcrumb-item>
            <a-breadcrumb-item>用户列表</a-breadcrumb-item>
          </a-breadcrumb>
        </div>
        <div class="rignt">
          <div class="actions">
            <div class="icon_actions">
              <i class="fa fa-home"></i>
              <GVBTheme></GVBTheme>

              <i class="fa fa-arrows-alt"></i>
            </div>
            <div class="avatar">
              <img src="https://img0.baidu.com/it/u=4054616240,2486283051&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500"
                alt="">
            </div>
            <div class="drop_menu">
              <a-dropdown>
                <a class="ant-dropdown-link" @click.prevent>
                  枫枫知道
                  <i class="fa fa-angle-down"></i>
                </a>
                <template #overlay>
                  <a-menu @click="menuDrop">
                    <a-menu-item key="user_center">
                      <a href="javascript:;">个人中心</a>
                    </a-menu-item>
                    <a-menu-item key="my_message">
                      <a href="javascript:;">我的消息</a>
                    </a-menu-item>
                    <a-menu-item key="article_list">
                      <a href="javascript:;">文章列表</a>
                    </a-menu-item>
                    <a-menu-item key="logout">
                      <a href="javascript:;">注销退出</a>
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </div>
          </div>

        </div>
      </header>
      <div class="tabs"></div>
      <main></main>
    </div>
  </div>
</template>

<script setup>
import GVBTheme from '../../components/gvb_theme.vue'
import GVBAside from '../../components/admin/gvb_aside.vue'
import { useRouter } from 'vue-router'
//下拉菜单点击事件
const router = useRouter()
const menuDrop = key  => {
  // console.log(key.key)
  if(key.key === 'logout'){
    return
  }
  router.push({name:key.key})
}


</script>

<style lang="scss">
.gvb_admin {
  width: 100%;
  display: flex;
}


.main {
  width: calc(100% - 240px);

  header {
    height: 60px;
    background-color: var(--hd);
    padding: 0 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .right {
      
    }
    .actions{
      align-items: center;
      display: flex;

    }
    .icon_actions {
      margin-right: 20px;

      i {
        margin-right: 10px;
        font-size: 20px;
        color: var(--text);
      }

      i:hover {
        color: var(--active);
      }
    }

    .avatar {
      img {
        width: 40px;
        height: 40px;
        border-radius: 50%;
      }
    }

    .drop_menu {
      margin-left: 10px;
    }
  }

  .tabs {
    height: 30px;
    border: 1px solid #f0eeee;
    background-color: var(--order);
  }

  main {
    background-color: var(--bg);
    height: calc(100vh - 90px);
  }
}
</style>