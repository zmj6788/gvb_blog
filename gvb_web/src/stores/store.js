import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useStore = defineStore('gvb',  {
  state: () => {
    return {
      theme:true
    }
  },
  actions: {
    // 设置主题
    setTheme(){
      this.theme = !this.theme
      if (this.theme) {
        //白色主题
        document.documentElement.classList.remove('dark')
        localStorage.setItem('theme',"light")
      } else {
        //黑色主题
        document.documentElement.classList.add('dark')
        localStorage.setItem('theme',"dark")
      }
    },
    // 加载主题
    loadTheme(){
      const theme = localStorage.getItem('theme')
      if(theme === "dark"){
        //黑色主题
        this.theme = false
        document.documentElement.classList.add('dark')
        return
      }
      this.theme = true

    }
  }
})
