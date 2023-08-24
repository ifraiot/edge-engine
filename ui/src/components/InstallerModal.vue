<template>
    <v-dialog v-model="showDialog" width="50%">
        <v-form validate-on="submit lazy" @submit.prevent="submit">
            <v-card>
                <v-card-text>
                    <v-select label="Application" item-title="label" item-value="id" v-model="selectedApplicationId"
                        :items="availableApplications" required></v-select>
                    <v-divider></v-divider>
                    <v-text-field v-for="(field, index) in selectedApplicationConfig" :key="index" :label="field.name"
                        v-model="selectedApplicationFormValues[field.id]" :hint="field.example"
                        :rules="field.is_required ? [v => !!v || `${field.name} field is required`] : []"
                        :v-if="selectedApplicationConfig != null && field.input_type !== 'volume-file'"></v-text-field>

                    <v-textarea label="Label"></v-textarea>
                </v-card-text>
                <v-card-actions>
                    <v-btn color="primary" type="submit">Start install</v-btn>
                    <v-btn color="info" @click="closeDialog">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-form>
    </v-dialog>
</template>

<script>
import axios from 'axios';
export default {
    props: {
        appType: {
            type: String
        },
        dialogVisible: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            availableApplications: [],
            selectedApplicationId: null,
            selectedApplicationConfig: [],
            selectedApplicationFormValues: {},
        }
    },
    computed: {
        showDialog: {
            get() {
                return this.dialogVisible;
            },
            set(value) {
                this.$emit('update:dialogVisible', value);
            },
        },
    },
    methods: {
        closeDialog() {
            this.$emit('update:dialogVisible', false)
        },
        async fetchAvailableApplicationsAPI(appTyp) {
            try {
                const response = await axios.get('http://localhost:8000/api/available-applications');
                const items = response.data.data;
                return items.filter(item => item.type === appTyp)
            } catch (error) {
                console.error('Error fetching items:', error);
            }
            return []
        },
    },
    watch: {
        async dialogVisible() {
            this.availableApplications = await this.fetchAvailableApplicationsAPI(this.appType)
        },
        selectedApplicationId(appId) {
            if (appId) {
                this.selectedApplicationFormValues = {}
                let selectApplication = this.availableApplications.find(item => item.id === appId);
                //Set Default Values    
                for (let index = 0; index < selectApplication.config.length; index++) {
                    const config = selectApplication.config[index];
                    this.selectedApplicationFormValues[config.id] = config.default === undefined ? '' : config.default
                }
                this.selectedApplicationConfig = selectApplication.config
            }
        }
    }
}
</script>

<style></style>