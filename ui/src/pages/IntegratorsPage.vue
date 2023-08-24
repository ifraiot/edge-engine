<template>
    <div>
        <v-row>
            <v-col cols="6" align-self="start">
                <h2>Integrator</h2>
            </v-col>
            <v-col cols="6" align-self="end" class="text-right">
                <v-btn class="pa-2 ma-2" color="primary" @click="addIntegratorHandler"> <v-icon>mdi-plus</v-icon> Install
                    Integrator</v-btn>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="12">
                <InstallerModal :dialogVisible="visibleInstallerModal"
                    @update:dialogVisible="visibleInstallerModal = $event" appType="integrator"></InstallerModal>
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
                        <tr v-for="item in installedApplications" :key="item.id">
                            <td>{{ item.application.label }}</td>
                            <td>{{ item.status }}</td>
                            <td>
                                <v-btn size="x-small">Log</v-btn>
                                <v-btn size="x-small">Edit</v-btn>
                                <v-btn @click="uninstallHandler(item.id)" size="x-small">Delete</v-btn>
                            </td>
                        </tr>
                    </tbody>
                </v-table>
            </v-col>
        </v-row>
    </div>
</template>
<script>
import axios from 'axios';
import InstallerModal from '../components/InstallerModal.vue'
export default {
    components: {
        InstallerModal
    },
    data() {
        return {
            installedApplications: [],
            visibleInstallerModal: false,
            loading: {
                delete: false
            }
        }
    },
    async created() {
        this.installedApplications = await this.fetchInstalledApplicationsAPI();
    },
    methods: {
        addIntegratorHandler() {
            this.visibleInstallerModal = true
        },
        async uninstallHandler(id) {
            this.loading.delete = true
            await this.uninstallApplicationAPI(id)
            this.installedApplications = await this.fetchInstalledApplicationsAPI();
        },
        async uninstallApplicationAPI(id) {
            try {
                const response = await axios.delete(`http://localhost:8000/api/applications/${id}`);
                console.log("API Response:", response.data);
                // Optionally, you can handle success or show a message to the user.
            } catch (error) {
                console.error("API Error:", error);
                // Handle error or show an error message to the user.
            }
        },
        async fetchInstalledApplicationsAPI() {
            try {
                var installedApplications = []
                const response = await axios.get('http://localhost:8000/api/applications');
                const items = response.data.data;
                for (let i = 0; i < items.length; i++) {
                    items[i].config = JSON.parse(items[i].config);
                    if (items[i].application.type === 'integrator') {
                        installedApplications.push(items[i]);
                    }
                }
                return installedApplications
            } catch (error) {
                console.error('Error fetching items:', error);
            }
            return []
        },
    },

  
}
</script>