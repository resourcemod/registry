<template>
  <default-layout>
    <div class="flex w-full h-full" v-if="content">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Delete {{content.type}}
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-2 gap-20 p-6 pt-2">
              <div class="flex flex-col gap-4 w-full">
                <div class="form-control">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Are you sure you want to delete the {{content.type}} <span class="font-semibold">{{content.name}}</span>? After that you will not be able to restore the record of this {{ content.type }}.</span>
                  </label>
                </div>

                <div class="form-control">
                  <button v-if="!isLoading" class="rounded-md py-2 px-6 bg-red-500 text-white" @click="deleteContent">Yes, delete</button>
                  <button v-else disabled="true" class="rounded-md py-2 px-6 bg-red-600 text-white">Deleting..</button>
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
import { format } from 'date-fns'
export default {
  data() {
    return {
      content: {
        type: '',
        name: '',
      },
      isLoading: false,
    }
  },
  methods: {
    async deleteContent() {
      this.isLoading = true
      try {
        this.errors = ''

        await this.$store.dispatch('deleteContent', {
          name: this.$route.params.name,
          type: this.$route.params.type,
        })

        location.href='/content/'+this.$route.params.type+'s'

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