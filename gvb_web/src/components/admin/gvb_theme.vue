<template>
  <i v-if="store.theme" class="fa fa-sun-o" @click="setTheme"></i>
  <i v-else class="fa fa-moon-o" @click="setTheme"></i>
</template>

<script setup>
import { useStore } from '@/stores/store'

//主题切换
const store = useStore()
const theme = store.theme //ture白色主题 false黑色主题
console.log(theme)
store.loadTheme()

//切换主题函数
const setTheme = () => {
  store.setTheme()
}
</script>

<style>

</style>