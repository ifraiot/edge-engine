 
<template>
  <v-card>
    <v-toolbar color="#00">
      <!-- <v-app-bar-nav-icon></v-app-bar-nav-icon> -->
      <v-toolbar-title color="#00">iFRA Edge Engine</v-toolbar-title>
      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>

      <template v-slot:extension>
        <v-tabs v-model="tab" align-tabs="title">
          <v-tab value="services">Services</v-tab>
          <v-tab value="two">Item Two</v-tab>
          <v-tab value="three">Item Three</v-tab>
        </v-tabs>

      </template>
    </v-toolbar>

    <v-window v-model="tab">
      <v-window-item key="services" value="services">
        <v-container>
          <v-row>
            <v-col>
              <!-- {{ installedApplications }} -->
              <h2>Connectors</h2>
              <v-btn @click="addServiceDrawer = true">Add</v-btn>
              <v-table>
                <thead>
                  <tr>
                    <th class="text-left">
                      Service
                    </th>
                    <th class="">
                      Status
                    </th>
                    <th class="text-left">
                      Manage
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in installedApplications.connectors" :key="item.app_id">
                    <td>{{ item.application.label }}</td>
                    <td>{{ item.status }}</td>
                    <td>
                      <v-btn @click="showConsolelog(item.id)" size="x-small">Log</v-btn>
                      <v-btn size="x-small">Edit</v-btn>
                      <v-btn :loading="loading.deleteLoading" @click="uninstallHandler(item.id)"
                        size="x-small">Delete</v-btn>
                    </td>
                  </tr>
                </tbody>
              </v-table>
            </v-col>
            <v-col>
              <h2>Analyzers</h2>
              <v-btn @click="addServiceDrawer = true">Add</v-btn>
              <v-table>
                <thead>
                  <tr>
                    <th class="text-left">
                      Service
                    </th>
                    <th class="text-left">
                      Status
                    </th>
                    <th class="text-left">
                      Manage
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in analyzers" :key="item.name">
                    <td>{{ item.name }}</td>
                    <td>{{ item.status }}</td>
                    <td><v-btn>Edit</v-btn></td>
                  </tr>
                </tbody>
              </v-table>
            </v-col>
            <v-col>
              <h2>Integrations</h2>
              <v-btn @click="addServiceDrawer = true">Add</v-btn>
              <v-table>
                <thead>
                  <tr>
                    <th class="text-left">
                      Service
                    </th>
                    <th class="text-left">
                      Status
                    </th>
                    <th class="text-left">
                      Manage
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in installedApplications.integrations" :key="item.app_id">
                    <td>{{ item.application.label }}</td>
                    <td>
                      {{ item.status }}
                    </td>
                    <td>
                      <v-btn @click="showConsolelog(item.id)" size="x-small">Log</v-btn>
                      <v-btn size="x-small">Edit</v-btn>
                      <v-btn :loading="loading.deleteLoading" @click="uninstallHandler(item.id)"
                        size="x-small">Delete</v-btn>
                    </td>
                  </tr>
                </tbody>
              </v-table>
            </v-col>
          </v-row>
          <v-row justify="space-around">
            <v-dialog v-model="addServiceDrawer" width="50%">
              <v-form validate-on="submit lazy" @submit.prevent="submit">
                <v-card>
                  <v-card-text>
                    <v-select label="Application" item-title="label" item-value="id"
                      v-model="selectedApplicationId" 
                      :items="availableApplications.integrations"
                      required></v-select>
                    <v-divider></v-divider>
                    <v-text-field :v-if="selectedApplicationConfig != null"
                      v-for="(field, index) in selectedApplicationConfig" :key="index" :label="field.name"
                      v-model="selectedApplicationFormValues[field.id]"
                      :hint="field.example"
                      :rules="field.is_required ? [v => !!v || `${field.name} field is required`] : []"></v-text-field>
                  </v-card-text>
                  <v-card-actions>
                    <v-btn color="primary" :loading="loading.installLoading" type="submit">Start install</v-btn>
                    <v-btn color="info" @click="addServiceDrawer = false">Close</v-btn>
                  </v-card-actions>
                </v-card>
              </v-form>
            </v-dialog>
          </v-row>
          <ConsoleLogVue logs="xxxxx" :visible="consoleLogVisible"></ConsoleLogVue>
        </v-container>
      </v-window-item>
    </v-window>
  </v-card>
</template>

<script>
import axios from 'axios';
import ConsoleLogVue from './ConsoleLog.vue';
export default {
  components: {
    ConsoleLogVue
  },
  data() {
    return {
      consoleLogVisible: false,
      loading: {
        deleteLoading: false,
        installLoading: false
      },
      selectedApplicationId: null,
      selectedApplication: null,
      selectedApplicationConfig: null,
      selectedApplicationFormValues: {},
      addServiceDrawer: false,
      tab: "services",
      connectors: [],
      analyzers: [],
      integrations: [],
      availableApplications: {
        connectors: [],
        analyzers: [],
        integrations: [],
      },
      installedApplications: {
        connectors: [],
        analyzers: [],
        integrations: [],
      }
    }
  },
  mounted() {
    this.fetchAvailableApplicationsWithAPI();
    this.fetchInstalledApplicationsWithAPI();
  },
  watch: {
    'selectedApplicationId': function (newAppId) {
      if (newAppId) {
        this.selectedApplicationFormValues = {}
        var connector = this.availableApplications.connectors.filter(item => item.id === newAppId);
        var analyzers = this.availableApplications.analyzers.filter(item => item.id === newAppId);
        var integrations = this.availableApplications.integrations.filter(item => item.id === newAppId);

        var selectApplication = null
        if (connector.length !== 0) {
          selectApplication = connector[0]
        } else if (analyzers.length !== 0) {
          selectApplication = analyzers[0]
        } else if (integrations.length !== 0) {
          selectApplication = integrations[0]
        }
        console.log(selectApplication)

        //Set Default Values
        for (let index = 0; index < selectApplication.config.length; index++) {
          const config = selectApplication.config[index];
          this.selectedApplicationFormValues[config.id] = config.default === undefined ? '' : config.default
        }



        this.selectedApplicationConfig = selectApplication.config
        console.log('Selected item:', selectApplication);
      }
    }
  },
  methods: {
    requiredRule(fieldId) {
      return [v => !!v || `${fieldId} is required`]; // Custom required validation rule
    },
    async submit(event) {
      const result = await event
      if(result.valid){
        this.installHandler()
      }
    },
    showConsolelog(id) {
      this.consoleLogVisible = true
      console.log(id)
    },
    async uninstallHandler(id) {
      this.loading.deleteLoading = true
      await this.uninstallApplicationWithAPI(id)
      this.loading.deleteLoading = false
      this.fetchInstalledApplicationsWithAPI();
    },
    async installHandler() {
      this.installLoading = true
      this.addServiceDrawer = false
      await this.installApplicationWithAPI()
      this.selectedApplicationId = null
      this.fetchInstalledApplicationsWithAPI();
      this.installLoading = false
    },
    async installApplicationWithAPI() {
      try {

        const keyValuePairs = Object.entries(this.selectedApplicationFormValues).map(([key, value]) => ({
          name: key,
          value: value,
        }));
        const response = await axios.post("http://localhost:8000/api/applications", {
          app_id: this.selectedApplicationId,
          configs: keyValuePairs
        });
        console.log("API Response:", response.data);
        // Optionally, you can handle success or show a message to the user.
      } catch (error) {
        console.error("API Error:", error);
        // Handle error or show an error message to the user.
      }
    },
    async uninstallApplicationWithAPI(id) {
      try {
        const response = await axios.delete(`http://localhost:8000/api/applications/${id}`);
        console.log("API Response:", response.data);
        // Optionally, you can handle success or show a message to the user.
      } catch (error) {
        console.error("API Error:", error);
        // Handle error or show an error message to the user.
      }
    },
    async fetchAvailableApplicationsWithAPI() {
      try {
        const response = await axios.get('http://localhost:8000/api/available-applications');
        const items = response.data.data;
        this.availableApplications.connectors = items.filter(item => item.type === 'connector');
        this.availableApplications.analyzers = items.filter(item => item.type === 'analyzer');
        this.availableApplications.integrations = items.filter(item => item.type === 'integrator');
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    },
    async fetchInstalledApplicationsWithAPI() {
      try {

        this.installedApplications.connectors = []
        this.installedApplications.analyzers = []
        this.installedApplications.integrations = []

        const response = await axios.get('http://localhost:8000/api/applications');
        const items = response.data.data;
        for (let i = 0; i < items.length; i++) {
          items[i].config = JSON.parse(items[i].config);
          if (items[i].application.type === 'connector') {
            this.installedApplications.connectors.push(items[i]);
          } else if (items[i].application.type === 'analyzer') {
            this.installedApplications.analyzers.push(items[i]);
          } else if (items[i].application.type === 'integrator') {
            this.installedApplications.integrations.push(items[i]);
          }
        }
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    },
  },
}
</script>
<style></style>