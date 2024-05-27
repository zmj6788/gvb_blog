<template>
  <aside class="gvb_aside">
    <div class="gvb_aside_header flex">
      <div class="gvb_aside_logo flex">
        <img src="https://img0.baidu.com/it/u=4054616240,2486283051&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500" alt="">
      </div>
      <div class="gvb_aside_title">
        枫枫知道博客后台
      </div>
    </div>
    <div class="gvb_aside_body flex">
      <div>
        <a-menu v-model:selectedKeys="state.selectedKeys" style="width: 256px; background-color: var(--slide);"
          mode="inline" :open-keys="state.openKeys" :items="items" @openChange="onOpenChange"></a-menu>
      </div>
    </div>
  </aside>

</template>

<script setup>
import { h, reactive } from 'vue';
import { MailOutlined, AppstoreOutlined, SettingOutlined } from '@ant-design/icons-vue';
function getItem(label, key, icon, children, type) {
  return {
    key,
    icon,
    children,
    label,
    type,
  };
}
const items = reactive([
  getItem('Navigation One', 'sub1', () => h(MailOutlined), [
    getItem('Option 1', '1'),
    getItem('Option 2', '2'),
    getItem('Option 3', '3'),
    getItem('Option 4', '4'),
  ]),
  getItem('Navigation Two', 'sub2', () => h(AppstoreOutlined), [
    getItem('Option 5', '5'),
    getItem('Option 6', '6'),
    getItem('Submenu', 'sub3', null, [getItem('Option 7', '7'), getItem('Option 8', '8')]),
  ]),
  getItem('Navigation Three', 'sub4', () => h(SettingOutlined), [
    getItem('Option 9', '9'),
    getItem('Option 10', '10'),
    getItem('Option 11', '11'),
    getItem('Option 12', '12'),
  ]),
]);
const state = reactive({
  rootSubmenuKeys: ['sub1', 'sub2', 'sub4'],
  openKeys: ['sub1'],
  selectedKeys: [],
});
const onOpenChange = openKeys => {
  const latestOpenKey = openKeys.find(key => state.openKeys.indexOf(key) === -1);
  if (state.rootSubmenuKeys.indexOf(latestOpenKey) === -1) {
    state.openKeys = openKeys;
  } else {
    state.openKeys = latestOpenKey ? [latestOpenKey] : [];
  }
};
</script>

<style lang="scss">
.gvb_aside {
  width: 240px;
  height: 100vh;
  background-color: var(--slide);

  .gvb_aside_header {
    width: 100%;
    height: 180px;
    flex-direction: column;
  }

  .gvb_aside_logo {
    padding: 10px;
    img{
      margin-top: 40px;
      width: 80px;
      height: 80px;
      border-radius: 10%;
    }
  }
  .gvb_aside_title {
    margin-top: 20px;
    font-size: 20px;
    color: #2c3e50;
    
    
  }
    .gvb_aside_body{
      width: 100%;
      margin-top: 50px;
    }

}
</style>