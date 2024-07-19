<template>
  <el-form ref="loginForm" :model="loginForm" label-width="80px">
    <el-form-item label="用户名">
      <el-input v-model="loginForm.username" autocomplete="off"></el-input>
    </el-form-item>
    <el-form-item label="密码">
      <el-input type="password" v-model="loginForm.password" autocomplete="off"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm('loginForm')">登录</el-button>
      <el-button @click="resetForm('loginForm')">重置</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { useAuthStore } from '@/stores/auth.js';
import { defineComponent } from 'vue';
import { useRouter } from 'vue-router';

export default defineComponent({
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      }
    };
  },
  methods: {
    async submitForm(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          try {
            const response = await this.$http.post('/user/login', {
              name: this.loginForm.username,
              password: this.loginForm.password,
            });
            if (response.data) {
              const authStore = useAuthStore();
              authStore.setToken(response.data.token, response.data.expire);
              alert('登录成功！');
              this.$router.push('/home');
            }
          } catch (error) {
            alert('登录失败，请检查用户名和密码！');
          }
        } else {
          alert('请输入正确的用户名和密码！');
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
});
</script>

<style>
.el-form {
  width: 240px;
  margin: 100px auto;
}
</style>
