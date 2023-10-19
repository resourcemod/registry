<template>
  <div class="h-screen w-screen flex flex-col justify-center items-center">
    <div class="flex flex-col">
      <div class="card flex-shrink-0 max-w-md bg-base-100">
        <div class="card-header text-center">
          <div>
            <img src="/assets/logo.png" alt="Registry logo" class="w-full" />
          </div>
          <div class="text-light-gray mt-2 mb-2">
            Sign up your new account
          </div>
        </div>
        <div class="flex flex-col gap-4 w-full">
          <div class="form-control">
            <label class="block text-sm text-font-gray">
              <span class="label-text">Name</span>
            </label>
            <input type="text" v-model="name" placeholder="" class="rounded w-full border-[1px] border-light-gray"/>
          </div>

          <div class="form-control">
            <label class="block text-sm text-font-gray">
              <span class="label-text">Password</span>
            </label>
            <input type="password" v-model="password" placeholder="" class="rounded w-full border-[1px] border-light-gray"/>
          </div>
          <div class="form-control">
            <label class="block text-sm text-font-gray">
              <span class="label-text">Confirm password</span>
            </label>
            <input type="password" v-model="password_confirmation" placeholder="" class="rounded w-full border-[1px] border-light-gray"/>
          </div>
          <div v-if="errors" class="text-error">
            {{errors}}
          </div>
          <div class="form-control flex flex-row justify-between items-center">
            <button class="rounded-md py-2 px-4 bg-font-gray text-white" @click="register">Sign Up</button>
            <router-link to="/register" class="rounded-md py-2 px-4 border-[1px] border-font-gray transition hover:bg-font-gray hover:text-white">Already have an account?</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>
<script setup lang="ts">
</script>

<script lang="ts">
export default {
  data() {
    return {
      name: '',
      password: '',
      password_confirmation: '',
      errors: ''
    }
  },
  methods: {
    validateName(name) {
      if (name.length > 253 || name.length <= 0) {
        return false
      }
      let reg = new RegExp(/[a-z][a-z0-9-.]{0,253}[a-z]$/, 'gm')
      return reg.test(name)
    },
    async register() {
      try {
        if (!this.validateName(this.name)) {
          this.errors = 'The name must comply with RFC 1123 Label Names standard'
          return
        }
        if (this.password != this.password_confirmation) {
          this.errors = 'The passwords must be same'
          return
        }
        const user = await this.$store.dispatch('register', {
          name: this.name,
          password: this.password,
          password_confirmation: this.password_confirmation
        })
        location.reload()
      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>