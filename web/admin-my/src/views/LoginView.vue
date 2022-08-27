<template>
<div class="container">
  <div class="loginBox">
    <a-form class="loginForm"
      :model="formState"
      :rules="rules"
    >
      <a-form-item
        label="Username"
        name="user"
      >
        <a-input
            v-model:value="formState.user"
            placeholder="username">
          <template #prefix><UserOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
        </a-input>
      </a-form-item>
      <a-form-item
          label="Password"
      >
        <a-input placeholder="password"
          type="password"
          v-model:value="formState.password"
        >
          <template #prefix><LockOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
        </a-input>
      </a-form-item>
      <a-form-item style="text-align: right">
        <a-button type="primary">登录</a-button>
        <a-button style="margin-left: 10px">取消</a-button>
      </a-form-item>
    </a-form>
  </div>
</div>
</template>

<script>
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
export default {
  name: "LoginView",
  components: {UserOutlined, LockOutlined},
  setup() {
    const formState = {
        user: "",
        password: ""
    }

    const rules = {
      user: [
        {require: true, message: "请输入用户名", trigger: blur},
        {min:4, max: 12, message: "用户名必须在4-12个字符之间", trigger:blur}
      ]
    }

    return {
      formState,
      rules
    }
  }

}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: bisque;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: lightyellow;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%,-50%);
  border-radius: 10px;

}
.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0px 20px;
  box-sizing: border-box;
}
</style>