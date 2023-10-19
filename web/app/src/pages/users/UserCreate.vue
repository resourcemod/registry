<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Create user
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-2 gap-20 p-6 pt-2">
              <div class="flex flex-col gap-4 w-full">
                <div v-if="errors" class="text-error">
                  {{errors}}
                </div>
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
                <div class="form-control mb-4">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Confirm password</span>
                  </label>
                  <input type="password" v-model="password_confirmation" placeholder="" class="rounded w-full border-[1px] border-light-gray"/>
                </div>

                <div class="form-control">
                  <button v-if="!isLoading" class="rounded-md py-2 px-6 bg-font-gray text-white" @click="createUser">Create</button>
                  <button v-else disabled="true" class="rounded-md py-2 px-6 bg-font-gray/80 text-white">Create</button>
                </div>
              </div>
              <div>

                <div class="relative w-72 lg:w-96 mb-4">
                  <input class="peer hidden" id="radio_1" type="radio" name="radio" v-model="role" v-bind:value="'user'" checked />
                  <span class="peer-checked:border-font-gray absolute right-4 top-1/2 box-content block h-3 w-3 -translate-y-1/2 rounded-full border-8 border-gray-300 bg-white"></span>
                  <label class="peer-checked:border-2 peer-checked:border-font-gray peer-checked:bg-brand-gray flex cursor-pointer select-none rounded-lg border border-gray-300 p-4 pr-20" for="radio_1">

                    <div class="ml-5">
                      <span class="mt-2 font-semibold">Regular user</span>
                      <p class="text-slate-500 text-sm leading-6">Can list and download content that have access to.</p>
                    </div>
                  </label>
                </div>
                <div class="relative w-72 lg:w-96">
                  <input class="peer hidden" id="radio_2" type="radio" name="radio" v-model="role" v-bind:value="'owner'" />
                  <span class="peer-checked:border-font-gray absolute right-4 top-1/2 box-content block h-3 w-3 -translate-y-1/2 rounded-full border-8 border-gray-300 bg-white"></span>

                  <label class="peer-checked:border-2 peer-checked:border-font-gray peer-checked:bg-brand-gray flex cursor-pointer select-none rounded-lg border border-gray-300 p-4 pr-20" for="radio_2">
                    <div class="ml-5">
                      <span class="mt-2 font-semibold">Registry Owner</span>
                      <p class="text-slate-500 text-sm leading-6">Can manage users, settings and any content.</p>
                    </div>
                  </label>
                </div>

              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </default-layout>
</template>
<script setup lang="ts">
import DefaultLayout from "../../components/layouts/DefaultLayout.vue";
</script>
<script lang="ts">
import {mapGetters} from "vuex";
import { format } from 'date-fns'
export default {
  data() {
    return {
      name: '',
      password: '',
      password_confirmation: '',
      errors: '',
      role: 'user',
      isLoading: false
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
    async createUser() {
      this.isLoading = true
      try {
        this.errors = ''
        if (!this.validateName(this.name)) {
          this.errors = 'The name must comply with RFC 1123 Label Names standard'
          return
        }
        if (this.$store.getters.getUserByName(this.name)) {
          this.errors = 'User with this name is already registered.'
          return
        }
        if (this.password != this.password_confirmation) {
          this.errors = 'The passwords must be same'
          return
        }

        const data = await this.$store.dispatch('createUser', {
          name: this.name,
          password: this.password,
          password_confirmation: this.password_confirmation,
          is_owner: this.role == 'owner'
        })

       location.href='/users'

      } catch (e) {
        this.errors = e
      }
      this.isLoading = false
    }
  },
  computed: {
    ...mapGetters(['getUser']),
  },
}
</script>