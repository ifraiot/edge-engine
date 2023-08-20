 
<template>
  <v-card>
    <v-toolbar color="primary">
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title>Edge Engine</v-toolbar-title>

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
                  <tr v-for="item in connectors" :key="item.name">
                    <td>{{ item.name }} {{ item.enabled }}</td>
                    <td>

                    </td>
                    <td><v-btn>Edit</v-btn></td>
                  </tr>
                </tbody>
              </v-table>
            </v-col>
            <v-col>
              <h2>Analyzers</h2>
              <v-btn>Add</v-btn>
              <v-table>
                <thead>
                  <tr>
                    <th class="text-left">
                      Name
                    </th>
                    <th class="text-left">
                      Calories
                    </th>
                    <th class="text-left">
                      Manage
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in analyzers" :key="item.name">
                    <td>{{ item.name }}</td>
                    <td>{{ item.calories }}</td>
                    <td><v-btn>Edit</v-btn></td>
                  </tr>
                </tbody>
              </v-table>
            </v-col>
            <v-col>
              <h2>Integrations</h2>
              <v-btn>Add</v-btn>
              <v-table>
                <thead>
                  <tr>
                    <th class="text-left">
                      Name
                    </th>
                    <th class="text-left">
                      Calories
                    </th>
                    <th class="text-left">
                      Manage
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in integrations" :key="item.name">
                    <td>{{ item.name }}</td>
                    <td>{{ item.calories }}</td>
                    <td><v-btn>Edit</v-btn></td>
                  </tr>
                </tbody>
              </v-table>
            </v-col>
          </v-row>
          <v-row justify="space-around">
            <v-dialog v-model="addServiceDrawer">
              <v-form validate-on="submit lazy" @submit.prevent="submit">
                <v-card>
                  <v-card-text>
                    <v-select label="Application" item-title="label" item-value="id" v-model="selectedApplicationId"
                      :items="availableApplications.connectors" required></v-select>
                    <v-divider></v-divider>
                    <v-text-field :v-if="selectedApplicationConfig != null"
                      v-for="(field, index) in selectedApplicationConfig" :key="index" :label="field.name"
                      v-model="selectedApplicationFormConfigs[field.id]"
                      :rules="field.is_required ? [v => !!v || `${field.name} field is required`] : []"></v-text-field>
                  </v-card-text>
                  <v-card-actions>
                    <v-btn color="primary" :loading="loading" type="submit">Save</v-btn>
                    <v-btn color="info" @click="addServiceDrawer = false">Close</v-btn>
                  </v-card-actions>
                </v-card>
              </v-form>
            </v-dialog>
          </v-row>
        </v-container>
      </v-window-item>
    </v-window>
  </v-card>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      loading: false,
      selectedApplication: null,
      selectedApplicationConfig: null,
      selectedApplicationId: null,
      selectedApplicationFormConfigs: {},
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
    }
  },
  mounted() {
    this.fetchAvailableApplications();
  },
  watch: {
    'selectedApplicationId': function (newAppId) {
      if (newAppId) {
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
      this.loading = true

      const results = await event

      this.loading = false

      alert(JSON.stringify(results))
      if (results.valid) {
        this.savedApplicationConfigHandler()
      }
    },
    async savedApplicationConfigHandler() {
      console.log(this.selectedApplicationFormValues)
      await this.installApplication()
      this.addServiceDrawer = false
      this.selectedApplicationId = null
    },
    async installApplication() {
      try {

        const keyValuePairs = Object.entries(this.selectedApplicationConfig).map(([key, value]) => ({
          name: key,
          value: value,
        }));
        const response = await axios.post("http://localhost:8000/api/install-application", {
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
    async fetchAvailableApplications() {
      try {
        const response = await axios.get('http://localhost:8000/api/available-applications');
        const items = response.data.data;
        this.availableApplications.connectors = items.filter(item => item.type === 'connector');
        this.availableApplications.analyzers = items.filter(item => item.type === 'analyzer');
        this.availableApplications.integrations = items.filter(item => item.type === 'integration');
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    },
  },
}
</script>
<style></style>