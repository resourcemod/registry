<template>
  <div class="text-moonlight h-screen flex flex-col justify-between bg-brand-gray border-r-[1px] border-dark-gray">
    <div>
      <div class="p-6 flex">
        <router-link to="/" class="block text-3xl font-semibold">registry.</router-link>
      </div>
      <div class="p-6">
        <nav>
          <div class="grid grid-cols-1 gap-1 pb-2">
          <router-link to="/" v-if="$route.name.startsWith('Dashboard')"
                       class="p-2 flex bg-font-gray rounded-md text-white">
          <span>
            <ChartBarIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Dashboard</span>
          </router-link>
          <router-link to="/" v-else class="transition p-2 flex hover:bg-font-gray rounded-md hover:text-white">
          <span>
            <ChartBarIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Dashboard</span>
          </router-link>

          <router-link to="/users" v-if="$route.name.startsWith('Users')"
                       class="p-2 flex bg-font-gray rounded-md text-white">
          <span>
            <UsersIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Users</span>
          </router-link>
          <router-link to="/users" v-else class="transition p-2 flex hover:bg-font-gray rounded-md hover:text-white">
          <span>
            <UsersIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Users</span>
          </router-link>

          <router-link to="/content/plugins" v-if="$route.name.startsWith('Plugins')"
                       class="p-2 flex bg-font-gray rounded-md text-white">
          <span>
            <CodeBracketIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Plugins</span>
          </router-link>
          <router-link to="/content/plugins" v-else class="transition p-2 flex hover:bg-font-gray rounded-md hover:text-white">
          <span>
            <CodeBracketIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Plugins</span>
          </router-link>


          <router-link to="/content/extension" v-if="$route.name.startsWith('Extensions')"
                       class="p-2 flex bg-font-gray rounded-md text-white">
          <span>
            <CircleStackIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Extensions</span>
          </router-link>
          <router-link to="/content/extensions" v-else class="transition p-2 flex hover:bg-font-gray rounded-md hover:text-white">
          <span>
            <CircleStackIcon class="h-4 w-4 mt-1"/>
          </span>
            <span class="ml-3">Extensions</span>
          </router-link>
</div>

          <div class="border-t-[1px] border-light-gray pt-2">
            <router-link to="/integrations" v-if="$route.name.startsWith('Integrations')"
                         class="p-2 flex bg-font-gray rounded-md text-white">
          <span>
            <BoltIcon class="h-4 w-4 mt-1"/>
          </span>
              <span class="ml-3">Integrations</span>
            </router-link>
            <router-link to="/integrations" v-else class="transition p-2 flex hover:bg-font-gray rounded-md hover:text-white">
          <span>
            <BoltIcon class="h-4 w-4 mt-1"/>
          </span>
              <span class="ml-3">Integrations</span>
            </router-link>
          </div>
          <div class="mt-1">
            <router-link to="/settings" v-if="$route.name.startsWith('Settings')"
                         class="p-2 flex bg-font-gray rounded-md text-white">
          <span>
            <Cog6ToothIcon class="h-4 w-4 mt-1"/>
          </span>
              <span class="ml-3">Settings</span>
            </router-link>
            <router-link to="/settings" v-else class="transition p-2 flex hover:bg-font-gray rounded-md hover:text-white">
          <span>
            <Cog6ToothIcon class="h-4 w-4 mt-1"/>
          </span>
              <span class="ml-3">Settings</span>
            </router-link>
          </div>
        </nav>
      </div>
    </div>
    <div class="p-6">
      <div class="flex items-center">
        <div>
          <div>
            Logged in as {{ getUser.name }}
          </div>
        </div>
      </div>
      <div class="text-sm flex justify-between mt-2">
        <div class="text-font-gray/30">RMod Registry v1.0.0</div>
        <router-link to="/logout" class="link">Logout</router-link>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {BoltIcon, ChartBarIcon, CircleStackIcon, CodeBracketIcon, UsersIcon, Cog6ToothIcon} from '@heroicons/vue/24/solid'
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