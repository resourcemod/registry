<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Upload content
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen">
            <div class="flex flex-row gap-4 p-6">
              <div class="relative w-72 lg:w-96">
                <input class="peer hidden" id="radio_1" type="radio" name="radio" v-model="type"
                       v-bind:value="'plugin'" checked/>
                <span
                    class="peer-checked:border-font-gray absolute right-4 top-1/2 box-content block h-3 w-3 -translate-y-1/2 rounded-full border-8 border-gray-300 bg-white"></span>
                <label
                    class="peer-checked:border-2 peer-checked:border-font-gray peer-checked:bg-brand-gray flex cursor-pointer select-none rounded-lg border border-gray-300 p-4 pr-20"
                    for="radio_1">

                  <div class="ml-5">
                    <span class="mt-2 font-semibold">Plugin</span>
                    <p class="text-slate-500 text-sm leading-6">JavaScript CS2 server modification that server owners
                      will install as plugin.</p>
                  </div>
                </label>
              </div>
              <div class="relative w-72 lg:w-96">
                <input class="peer hidden" id="radio_2" type="radio" name="radio" v-model="type"
                       v-bind:value="'extension'"/>
                <span
                    class="peer-checked:border-font-gray absolute right-4 top-1/2 box-content block h-3 w-3 -translate-y-1/2 rounded-full border-8 border-gray-300 bg-white"></span>

                <label
                    class="peer-checked:border-2 peer-checked:border-font-gray peer-checked:bg-brand-gray flex cursor-pointer select-none rounded-lg border border-gray-300 p-4 pr-20"
                    for="radio_2">
                  <div class="ml-5">
                    <span class="mt-2 font-semibold">Extension</span>
                    <p class="text-slate-500 text-sm leading-6">Native QuickJS C/C++ module to help ResourceMod works
                      faster.</p>
                  </div>
                </label>
              </div>
            </div>
            <div class="w-full p-6 pt-2">
              <div class="grid grid-cols-2 gap-4 w-full">
                <div class="form-control col-span-2 w-[300px]">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Name</span>
                  </label>
                  <input type="text" v-model="name" placeholder=""
                         class="rounded w-full border-[1px] border-light-gray"/>
                </div>

                <div class="form-control col-span-2 w-[300px]">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Version</span>
                  </label>
                  <input type="text" v-model="version" placeholder=""
                         class="rounded w-full border-[1px] border-light-gray"/>
                </div>

                <div class="form-control col-span-2">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Access</span>
                  </label>
                  <select class="rounded w-[300px] border-[1px] border-light-gray" v-model="is_public">
                    <option value="1">Public</option>
                    <option value="0">Private</option>
                  </select>
                </div>

                <div class="form-control col-span-2">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Description (plain text)</span>
                  </label>
                  <textarea class="rounded w-full border-[1px] border-light-gray" v-model="description" cols="30" rows="10"></textarea>
                </div>
                <div v-if="errors" class="text-error">
                  {{ errors }}
                </div>
                <div class="form-control">
                  <button v-if="!isLoading" class="rounded-md py-2 px-6 bg-font-gray text-white" @click="uploadContent">
                    Create
                  </button>
                  <button v-else disabled="true" class="rounded-md py-2 px-6 bg-font-gray/80 text-white">Uploading..
                  </button>
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
import DefaultLayout from "../components/layouts/DefaultLayout.vue";
</script>
<script lang="ts">
import {mapGetters} from "vuex";

export default {
  data() {
    return {
      name: '',
      type: 'plugin',
      version: '1.0',
      is_public: 1,
      description: '',
      errors: '',
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
    async addPlugin() {
      this.isLoading = true
      try {
        this.errors = ''
        if (!this.validateName(this.name)) {
          this.errors = 'The name must comply with RFC 1123 Label Names standard'
          return
        }
        if (this.$store.getters.getPluginByName(this.name)) {
          this.errors = 'User with this name is already registered.'
          return
        }

        const data = await this.$store.dispatch('createContent', {
          type: 'plugin',
          name: this.name,

        })

        location.href = '/content/plugins'

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

<style>
input:checked ~ span:last-child {

  --tw-translate-x: 3.54rem; /* translate-x-7 */
}
h1 {
  font-weight: bold;
  font-size: 36px;
}
h2 {
  font-weight: semi-bold;
  font-size: 26px;
}
h3 {
  font-weight: medium;
  font-size: 18px;
}
h4 {
  font-size: 16px;
}

ul {
  list-style-type: circle;
}
</style>