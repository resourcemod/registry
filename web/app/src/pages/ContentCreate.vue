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

                <div class="form-control col-span-2">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Git integration</span>
                  </label>
                  <select @change="updateRepositories" class="rounded w-[300px] border-[1px] border-light-gray" v-model="selected_integration">
                    <option v-for="integration in integrations" :value="integration.name">{{ integration.name }}</option>
                  </select>
                </div>


                <div class="form-control col-span-2" v-if="this.repositories">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Repository</span>
                  </label>
                  <select class="rounded w-[300px] border-[1px] border-light-gray" v-model="selected_repository">
                    <option v-for="repository in repositories" :value="repository">{{ repository.full_name }}</option>
                  </select>
                </div>

                <div v-if="errors" class="text-error col-span-2">
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
      repositories: [],
      integrations: [],
      selected_repository: '',
      selected_integration: '',
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
    async updateRepositories() {
      if (this.selected_integration !== '') {
        const data = await this.$store.dispatch('getRepositories', {name: this.selected_integration})
        this.repositories = data.repositories
      }
    },
    async uploadContent() {
      this.isLoading = true
      try {
        this.errors = ''
        if (!this.validateName(this.name)) {
          this.errors = 'The name must comply with RFC 1123 Label Names standard'
          this.isLoading = false
          return
        }
        if (this.$store.getters.getContentByName(this.type, this.name)) {
          this.errors = 'Name is already taken.'
          this.isLoading = false
          return
        }

        if (this.selected_repository === '') {
          this.errors = 'Repository is required.'
          this.isLoading = false
          return
        }

        await this.$store.dispatch('uploadContent', {
          name: this.name,
          description: this.description,
          version: this.version,
          type: this.type,
          is_public: !!this.is_public,
          repository: this.selected_repository
        })

        location.href = '/content/'+this.type+'s'

      } catch (e) {
        this.errors = e.message
      }
      this.isLoading = false
    }
  },
  async mounted() {
    if (this.integrations.length === 0) {
      const data  = await this.$store.dispatch('getIntegrations')
      this.integrations = data.integrations
    }
    if (this.selected_integration !== '') {
      this.repositories = await this.$store.dispatch('getRepositories', this.selected_integration)
    }
    if (this.$route.query.type == 'extension') {
      this.type = 'extension'
    }
  },
  computed: {
    ...mapGetters(['getUser', 'getIntegrations']),
  },
}
</script>