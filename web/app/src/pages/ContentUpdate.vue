<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Update content
          </div>
          <div>
            <router-link :to="'/content/'+content.type+'/'+content.name+'/delete'" class="bg-font-gray rounded text-white py-2 px-4 mr-2 lowercase text-sm"><MinusIcon class="w-4 h-4 inline -mt-[1px]"/> delete</router-link>
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen" v-if="content">
            <div class="flex flex-row gap-4 p-6">
              <div class="relative w-72 lg:w-96">
                <input disabled class="peer hidden" id="radio_1" type="radio" name="radio" v-model="content.type"
                       v-bind:value="'plugin'"/>
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
                <input disabled class="peer hidden" id="radio_2" type="radio" name="radio" v-model="content.type"
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
                    <span class="label-text">Name (read-only)</span>
                  </label>
                  <input disabled type="text" v-model="content.name" placeholder=""
                         class="rounded w-full border-[1px] border-light-gray"/>
                </div>

                <div class="form-control w-[300px]">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Version (change it to create a new revision)</span>
                  </label>
                  <input type="text" v-model="content.version" placeholder=""
                         class="rounded w-full border-[1px] border-light-gray"/>
                  <div class="text-gray-400 text-sm">
                    <div class="label-text">Current version: {{content.release.version}}</div>
                    <div class="label-text">Current git release: {{content.release.release_name}}</div>
                  </div>
                </div>

                <div class="form-control col-span-2">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Access</span>
                  </label>
                  <select class="rounded w-[300px] border-[1px] border-light-gray" v-model="content.is_public">
                    <option :value="true">Public</option>
                    <option :value="false">Private</option>
                  </select>
                </div>

                <div class="form-control col-span-2">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Description (plain text)</span>
                  </label>
                  <textarea class="rounded w-full border-[1px] border-light-gray" v-model="content.description" cols="30" rows="10"></textarea>
                </div>

                <div v-if="errors" class="col-span-2 text-error">
                  {{ errors }}
                </div>
                <div class="form-control">
                  <button v-if="!isLoading" class="rounded-md py-2 px-6 bg-font-gray text-white" @click="uploadContent">
                    Update
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
import {MinusIcon} from "@heroicons/vue/24/solid";
</script>
<script lang="ts">
import {mapGetters} from "vuex";

export default {
  data() {
    return {
      content: {
        name: '',
        type: 'plugin',
        version: '1.0',
        is_public: 1,
        description: '',
        release: {
          release_name: '',
          version: '',
        }
      },
      errors: '',
      isLoading: false
    }
  },
  methods: {
    async uploadContent() {
      this.isLoading = true
      try {
        this.errors = ''

        await this.$store.dispatch('updateContent', {
          type: this.content.type,
          name: this.content.name,
          description: this.content.description,
          version: this.content.version,
          is_public: !!this.content.is_public,
        })

      } catch (e) {
        this.errors = e.message
      }
      this.isLoading = false
    }
  },
  async mounted() {
    this.content = await this.$store.dispatch('getContentByTypeAndName', {type: this.$route.params.type, name: this.$route.params.name})
  },
}
</script>