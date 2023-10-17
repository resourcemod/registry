<template>
  <div class="min-h-screen mx-auto">
    <div class="flex-col lg:flex-row-reverse">
      <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100 rounded-xl">
        <div class="card-header text-center">
          <h1 class="h1 text-2xl pt-4">ResourceMod Registry</h1>
          <div class="text-sm">
            Create owner account
          </div>
        </div>
        <div class="card-body w-[370px]">
          <div class="form-control">
            <label class="label">
              <span class="label-text">Name</span>
            </label>
            <input type="text" v-model="name" placeholder="" class="input input-bordered"/>
          </div>

          <div class="form-control">
            <label class="label">
              <span class="label-text">Password</span>
            </label>
            <input type="password" v-model="password" placeholder="" class="input input-bordered"/>
          </div>
          <div class="form-control">
            <label class="label">
              <span class="label-text">Confirm password</span>
            </label>
            <input type="password" v-model="password_confirmation" placeholder="" class="input input-bordered"/>
          </div>
          <div v-if="errors" class="text-error">
            {{errors}}
          </div>
          <div class="form-control mt-6">
            <button class="rounded-xl btn btn-primary" @click="createFirstAccount">Create an account</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>
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
    async createFirstAccount() {
      try {
        if (!this.validateName(this.name)) {
          this.errors = 'The name must comply with RFC 1123 Label Names standard'
          return
        }
        if (this.password != this.password_confirmation) {
          this.errors = 'The passwords must be same'
          return
        }

        const data = await this.$store.dispatch('createFirstAccount', {
          name: this.name,
          password: this.password,
          password_confirmation: this.password_confirmation
        })

      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>