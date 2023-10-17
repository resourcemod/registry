<template>
  <div class="text-moonlight h-screen flex flex-col justify-between">
    <div>
      <div class="p-4 flex">
        <span class="block text-2xl font-semibold uppercase">resourcemod registry</span>
      </div>
      <div class="p-4">
        <nav class="grid grid-cols-1 gap-1">
          <router-link to="/" v-if="$route.name.startsWith('Dashboard')"
                       class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <ChartBarIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Dashboard</span>
          </router-link>
          <router-link to="/" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <ChartBarIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Dashboard</span>
          </router-link>

          <router-link to="/users" v-if="$route.name.startsWith('Users')"
                       class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <UsersIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Users</span>
          </router-link>
          <router-link to="/users" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <UsersIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Users</span>
          </router-link>

          <router-link to="/content/plugin" v-if="$route.name.startsWith('Plugins')"
                       class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <ServerStackIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Plugins</span>
          </router-link>
          <router-link to="/content/plugin" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <ServerStackIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Plugins</span>
          </router-link>

          <router-link to="/content/extension" v-if="$route.name.startsWith('Extensions')"
                       class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <CubeIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Extensions</span>
          </router-link>
          <router-link to="/content/extension" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <CubeIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Extensions</span>
          </router-link>

          <div class="pt-4 border-t-[1px] border-white/10">
            <router-link to="/integrations" v-if="$route.name.startsWith('Integrations')"
                         class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <BoltIcon class="h-4 w-4 mt-1"/>
          </span>
              <span class="ml-3">Integrations</span>
            </router-link>
            <router-link to="/integrations" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <BoltIcon class="h-4 w-4 mt-1"/>
          </span>
              <span class="ml-3">Integrations</span>
            </router-link>
          </div>
        </nav>
      </div>
    </div>
    <div class="p-4">
      <div class="flex items-center p-4 hover:border-white/20 rounded-md border-[1px] border-white/10">
        <div>
          <div>
            {{ getUser.name }}
          </div>
        </div>
      </div>
      <div class="text-sm flex justify-between mt-2">
        <div class="text-white/30">RMod Registry v1.0.0</div>
        <router-link to="/logout" class="link">Logout</router-link>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {BoltIcon, ChartBarIcon, CubeIcon, ServerStackIcon, UsersIcon} from '@heroicons/vue/24/solid'
import {mapGetters} from 'vuex'
</script>

<script lang="ts">
import {mapGetters} from "vuex";

export default {
  computed: {
    ...mapGetters(['getUser'])
  }
}
</script>