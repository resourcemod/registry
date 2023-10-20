<template>
  <default-layout>
    <div class="flex w-full overflow-y-auto">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Plugins
          </div>
          <div>
            <router-link to="/content/create" class="bg-font-gray rounded text-white py-2 px-4 mr-2 lowercase text-sm"><PlusIcon class="w-4 h-4 inline -mt-[1px]"/> Upload</router-link>
          </div>
        </div>
        <div class="pl-4" v-if="getPlugins">
          <div class="">
            <div class="w-full p-6">
              <div v-if="getPlugins.size == 0">
                No plugins found.
              </div>
              <div v-else>
                asdasd
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