<template>
  <default-layout>
    <div class="flex w-full items-start">
      <div class="flex flex-col w-full gap-4">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Git Integrations
          </div>
        </div>
        <div class="pl-4">
          <div class="overflow-y-auto">
            <div class="w-full grid grid-cols-3 gap-6 p-4 py-0">
              <div class="w-full flex flex-col gap-4 rounded-md border-[1px] border-gray-200 p-4">
                <div class="form-control">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Host</span>
                  </label>
                  <input type="text" disabled v-model="host" placeholder=""
                         class="text-gray-400 rounded bg-blue-50 w-full border-[1px] border-light-gray"/>
                </div>

                <div class="form-control">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Name</span>
                  </label>
                  <input type="text" v-model="name" placeholder=""
                         class="rounded w-full border-[1px] border-light-gray"/>
                </div>

                <div class="form-control">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Auth Token</span>
                  </label>
                  <input type="text" v-model="access_token" placeholder=""
                         class="rounded w-full border-[1px] border-light-gray"/>
                </div>

                <div class="form-control ">
                  <button v-if="!isLoading" class="w-full rounded-md py-2 px-6 bg-font-gray text-white" @click="connectIntegration">
                    Connect
                  </button>
                  <button v-else disabled="true" class="w-full rounded-md py-2 px-6 bg-font-gray/80 text-white">Connecting..
                  </button>
                </div>
              </div>


              <div class="w-full flex flex-col gap-4 rounded-md border-[1px] border-gray-200 p-4" v-for="[key, integration] in getIntegrations">
                <div class="form-control">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Host</span>
                  </label>
                  <input type="text" disabled v-model="integration.host" placeholder=""
                         class="text-gray-400 rounded bg-blue-50 w-full border-[1px] border-light-gray"/>
                </div>

                <div class="form-control">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Name</span>
                  </label>
                  <input type="text" disabled v-model="integration.name" placeholder=""
                         class="text-gray-400 rounded w-full bg-blue-50 border-[1px] border-light-gray"/>
                </div>

                <div class="form-control">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Auth Token</span>
                  </label>
                  <input type="text" disabled value="hidden." placeholder=""
                         class="italic text-gray-400 rounded w-full bg-blue-50 border-[1px] border-light-gray"/>
                </div>

                <div class="form-control ">
                  <button v-if="!isDeleting[integration.name]" class="w-full rounded-md py-2 px-6 bg-red-500 text-white" @click="deleteIntegration(integration)">
                    Delete
                  </button>
                  <button v-else disabled="true" class="w-full rounded-md py-2 px-6 bg-font-gray/80 text-white">Deleting..
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
import DefaultLayout from "../../components/layouts/DefaultLayout.vue";
import {UserPlusIcon, ListBulletIcon} from '@heroicons/vue/24/solid'
</script>

<script lang="ts">
import {formatDistance} from "date-fns";
import {mapGetters} from "vuex";
import { useToast } from "vue-toastification";
const toast = useToast();

export default {
  data() {
    return {
      name: '',
      type: 'git',
      host: 'github.com',
      access_token: '',
      isLoading: false,
      isDeleting: [],
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
    async deleteIntegration(integration) {
      this.isDeleting[integration.name] = true
      try {
        const data = await this.$store.dispatch('deleteIntegration', {
          name: integration.name,
        })

      } catch (e) {
        toast.error(e)
      }
      this.isDeleting[integration.name] = false
    },
    async connectIntegration() {
      this.isLoading = true
      try {
        if (!this.validateName(this.name)) {
          toast.error('The name must comply with RFC 1123 Label Names standard')
          this.isLoading = false
          return
        }
        if (this.name == "") {
          toast.error('Name is required')
          this.isLoading = false
          return
        }
        if (this.access_token == "") {
          toast.error('Auth token is required')
          this.isLoading = false
          return
        }

        const data = await this.$store.dispatch('createIntegration', {
          name: this.name,
          host: this.host,
          type: this.type,
          access_token: this.access_token
        })
        this.access_token = ''
        this.name = ''

        toast.success("Git integration was connected")

        await this.$store.dispatch('getIntegrations')

      } catch (e) {
        toast.error(e)
      }
      this.isLoading = false
    }
  },
  async mounted() {
    if (this.getIntegrations.size === 0) {
      await this.$store.dispatch('getIntegrations')
    }
  },
  computed: {
    ...mapGetters(['getIntegrations'])
  }
}
</script>