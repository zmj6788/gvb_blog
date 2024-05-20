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
    <div class="gvb_aside_body">
      
    </div>
  </aside>

</template>

<script setup>

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

}
</style>