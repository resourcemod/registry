<template>
  <default-layout>
    <div class="flex w-full items-start">
      <div class="flex flex-col w-full gap-4">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Users
          </div>
          <div>
            <router-link to="/users/feed" class="bg-font-gray rounded text-white py-2 px-4 mr-2 lowercase text-sm"><ListBulletIcon class="w-4 h-4 inline -mt-[1px]"/> Activity Feed</router-link>
            <router-link to="/users/create" class="bg-font-gray rounded text-white py-2 px-4 mr-2 lowercase text-sm"><UserPlusIcon class="w-4 h-4 inline -mt-[1px]"/> Create user</router-link>
          </div>
        </div>
        <div class="p-6">
          <div class="overflow-y-auto">
            <div class="w-full">
              <table class="min-w-full border-collapse border-spacing-y-2 border-spacing-x-2">
                <thead class="hidden border-b lg:table-header-group">
                <tr class="">
                  <td class="whitespace-normal py-4 text-sm font-semibold text-gray-800 sm:px-3">
                    Name
                  </td>
                  <td class="whitespace-normal py-4 text-sm font-medium text-gray-500 sm:px-3">Created</td>
                  <td class="whitespace-normal py-4 text-sm font-medium text-gray-500 sm:px-3">Updated</td>

                  <td class="whitespace-normal py-4 text-sm font-medium text-gray-500 sm:px-3">Role</td>
                  <td class="whitespace-normal py-4 text-sm font-medium text-gray-500 sm:px-3"></td>
                </tr>
                </thead>

                <tbody class="bg-white lg:border-gray-300">
                <tr class="" v-for="[key, user] in getUsers">
                  <td class="whitespace-no-wrap py-4 text-left text-sm font-medium sm:px-3 lg:text-left">
                    {{ user.name }}
                  </td>
                  <td class="whitespace-no-wrap hidden py-4 text-sm font-normal sm:px-3 lg:table-cell">{{ formatDistance(new Date(user.created_at), new Date()) }}</td>
                  <td class="whitespace-no-wrap hidden py-4 text-sm font-normal sm:px-3 lg:table-cell">{{ formatDistance(new Date(user.updated_at), new Date()) }}</td>
                  <td class="whitespace-no-wrap hidden py-4 text-sm font-normal sm:px-3 lg:table-cell">
                    <span v-if="user.is_owner" class="whitespace-nowrap rounded-full bg-green-100 px-2 py-0.5 text-green-800">owner</span>
                    <span v-else class="whitespace-nowrap rounded-full bg-gray-100 px-2 py-0.5 text-green-800">user</span>
                  </td>
                  <td class="whitespace-no-wrap hidden py-4 text-sm font-normal text-gray-500 sm:px-3 lg:table-cell">
                    <router-link :to="'/users/'+user.name" class="ml-2 mr-3 rounded-md border-[1px] border-font-gray px-4 py-2 text-font-gray hover:text-white hover:bg-font-gray transition">details</router-link>
                  </td>
                </tr>
                </tbody>
              </table>

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

export default {
  async mounted() {
    if (this.getUsers.size === 0) {
      await this.$store.dispatch('getUsers')
    }
  },
  computed: {
    ...mapGetters(['getUsers'])
  }
}
</script>