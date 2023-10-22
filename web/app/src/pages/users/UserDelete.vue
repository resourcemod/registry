<template>
  <default-layout>
    <div class="flex w-full h-full" v-if="user">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Delete user
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-2 gap-20 p-6 pt-2">
              <div class="flex flex-col gap-4 w-full">
                <div class="form-control">
                  <label class="block text-sm text-font-gray">
                    <span class="label-text">Are you sure you want to delete the user <span class="font-semibold">{{user.name}}</span>? After that you will not be able to restore the record of this user.</span>
                  </label>
                </div>

                <div class="form-control">
                  <button v-if="!isLoading" class="rounded-md py-2 px-6 bg-red-500 text-white" @click="deleteUser">Yes, delete</button>
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
import DefaultLayout from "../../components/layouts/DefaultLayout.vue";
</script>
<script lang="ts">
import {mapGetters} from "vuex";
import { format } from 'date-fns'
export default {
  data() {
    return {
      user: null,
      isLoading: false,
    }
  },
  methods: {
    async deleteUser() {
      this.isLoading = true
      try {
        this.errors = ''

        const data = await this.$store.dispatch('deleteUser', {
          name: this.user.name,
        })

       location.href='/users'

      } catch (e) {
        this.errors = e
      }
      this.isLoading = false
    }
  },
  computed: {
    ...mapGetters(['getUsers']),
  },
  async mounted() {
    if (this.getUsers.size === 0) {
      await this.$store.dispatch('getUsers')
    }
    this.user = this.$store.getters.getUserByName(this.$route.params.name)
  }
}
</script>