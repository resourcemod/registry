<template>
  <default-layout>
    <div class="flex w-full">
      <div class="flex flex-col w-full"  v-if="user">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
            <div class="text-xl pl-10 py-4">
              User details
            </div>
            <div>
              <router-link :to="'/users/'+user.name+'/delete'" class="cursor-pointer bg-red-500 rounded text-white py-2 px-4 mr-2 lowercase text-sm"><UserMinusIcon class="w-4 h-4 inline -mt-[1px]"/> delete</router-link>
            </div>
        </div>
        <div class="pl-4 mt-4">
          <div>
            <div class="w-full flex flex-row gap-20 p-6 pt-2">
              <div class="w-1/4 flex flex-col gap-1">
                <div class="font-semibold text-lg">About</div>
                <!-- About Section -->
                <div class="font-medium">Name: {{user.name}}</div>
                <div>Role:
                  <span v-if="user.is_owner" class="whitespace-nowrap rounded-full bg-green-100 px-2 py-0.5 text-green-800">owner</span>
                  <span v-else class="whitespace-nowrap rounded-full bg-gray-100 px-2 py-0.5 text-green-800">user</span>
                </div>
                <div>Created: {{formatDistance(new Date(user.created_at), new Date())}}</div>
                <div>Updated: {{formatDistance(new Date(user.updated_at), new Date())}}</div>
                <!-- End of about section -->
              </div>
              <div class="w-full">
                <div class="font-semibold text-lg">Activity feed</div>
                <div>

                  <ul aria-label="User feed" role="feed" class="relative flex flex-col gap-6 py-12 before:absolute before:top-0 before:h-full before:border before:-translate-x-1/2 before:border-slate-200 before:border-dashed after:absolute after:top-6 after:bottom-6 after:border after:-translate-x-1/2 after:border-slate-200 ">
                    <li role="article" class="relative pl-8 ">
                      <div class="flex flex-col flex-1 gap-4">
                        <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin_2: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 1 hour ago</span></h4>
                      </div>
                    </li>
                    <li role="article" class="relative pl-8 ">
                      <div class="flex flex-col flex-1 gap-4">
                        <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                      </div>
                    </li>







                    <li role="article" class="relative pl-8 ">
                      <div class="flex flex-col flex-1 gap-4">
                        <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                      </div>
                    </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li><li role="article" class="relative pl-8 ">
                    <div class="flex flex-col flex-1 gap-4">
                      <h4 class="flex flex-col items-start text-lg font-medium leading-8 lg:items-center md:flex-row text-slate-700"><span class="text-base font-normal text-slate-500">Installed :example_plugin: via package manager</span><span class="text-sm font-normal text-slate-400 ml-4"> 3 hours ago</span></h4>
                    </div>
                  </li>





                  </ul>


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
import {formatDistance} from "date-fns";
import {UserMinusIcon} from "@heroicons/vue/24/solid";


import { ref } from 'vue'
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  DialogDescription,
} from '@headlessui/vue'

const isOpen = ref(true)

function setIsOpen(value) {
  isOpen.value = value
}

</script>
<script lang="ts">
import {mapGetters} from "vuex";
import { format } from 'date-fns'
export default {
  data() {
    return {
      user: null
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