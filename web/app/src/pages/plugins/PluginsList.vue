<template>
  <default-layout>
    <div class="flex w-full overflow-y-auto">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Plugins
          </div>
          <div>
            <router-link to="/content/create?type=plugin" class="bg-font-gray rounded text-white py-2 px-4 mr-2 lowercase text-sm"><PlusIcon class="w-4 h-4 inline -mt-[1px]"/> Upload</router-link>
          </div>
        </div>
        <div class="pl-4" v-if="getPlugins">
          <div class="">
            <div class="w-full p-6">
              <div v-if="getPlugins.size == 0">
                No plugins found.
              </div>
              <div v-else>
                <div class="grid grid-cols-3 gap-6">
                  <div class="border-[1px] rounded-lg p-4" v-for="[key, content] in getPlugins">
                    <div class="text-medium truncate"><span class="bg-gray-200 px-2 py-1 text-sm mr-2 rounded">{{content.is_public ? 'Public': 'Private'}}</span>{{content.name}}</div>
                    <div class="text-medium truncate text-gray-400">{{content.description}}</div>
                    <div class="text-medium w-full truncate">Created {{formatDistance(new Date(content.created_at), new Date())}}</div>
                    <div class="text-medium w-full truncate">Updated {{formatDistance(new Date(content.updated_at), new Date())}}</div>
                    <div class="text-medium w-full truncate mt-4">
                      <router-link :to="'/content/plugin/'+content.name" class="block bg-font-gray rounded text-white cursor-pointer text-center px-4 py-2" >Update</router-link>
                    </div>
                  </div>
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
import {PlusIcon} from "@heroicons/vue/24/solid";
import {formatDistance} from "date-fns";
</script>

<script lang="ts">
import {mapGetters} from "vuex";

export default {
  async mounted() {
    if (this.getPlugins.size === 0) {
      await this.$store.dispatch('getContentByType', {
        type: 'plugin'
      })
    }
  },
  computed: {
    ...mapGetters(['getPlugins'])
  }
}
</script>