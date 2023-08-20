 
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
              <v-card>
 
                <v-card-text>
                 
                  <v-select label="Select" item-title="Name" item-value="Name" :items="availableApplications.connectors"></v-select>

                  <v-text-field label="Phone Number"></v-text-field>

                  <v-text-field label="E-mail"></v-text-field>

                  <v-select label="Select"></v-select>

                  <v-checkbox label="Option" type="checkbox"></v-checkbox>

                  <v-btn class="me-4" type="submit">
                    submit
                  </v-btn>

                  <v-btn>
                    clear
                  </v-btn>
                </v-card-text>
                <v-card-actions>
                  <v-btn color="primary" block @click="dialog = false">Close Dialog</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-row>

          <!-- <v-navigation-drawer location="right" v-model="addServiceDrawer"> -->

          <!-- <form>
             
            </form> -->
          <!-- <v-btn color="primary" @click.stop="addServiceDrawer = !addServiceDrawer">
              Toggle
            </v-btn> -->

          <!-- </v-navigation-drawer> -->


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
      addServiceDrawer: false,
      tab: "services",
      connectors: [],
      analyzers: [],
      integrations: [],
      availableApplications: {
        connectors: [],
        analyzers: [],
        integrations: [],
      }
    }
  },
  mounted() {
    // Make the API call when the component is mounted
    this.fetchAvailableApplications();
  },
  methods: {
    async fetchAvailableApplications() {
      try {
        const response = await axios.get('http://localhost:8000/api/available-applications');
        const items = response.data.data;
        this.availableApplications.connectors = items.filter(item => item.Type === 'connector');
        this.availableApplications.analyzers = items.filter(item => item.Type === 'analyzer');
        this.availableApplications.integrations = items.filter(item => item.Type === 'integration');
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    },
  },
}
</script>
<style></style>