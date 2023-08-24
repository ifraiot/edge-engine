import { createApp } from 'vue'
import App from './App.vue'

import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { fa } from "vuetify/iconsets/fa";
import { aliases, mdi } from "vuetify/lib/iconsets/mdi";
import "@mdi/font/css/materialdesignicons.css";
import "@fortawesome/fontawesome-free/css/all.css";
import router from './router/router';
const myAllBlackTheme = {
  dark: false,
  colors: {
    background: "#FFFF",
    surface: "#FFFF",
    primary: "#0097a7ff",
    'primary-darken-1': '#3700B3',
    secondary: '#03DAC6',
    'secondary-darken-1': '#018786',
    error: "#000000",
    info: "#000000",
    success: "#000000",
    warning: "#000000",
  }
};


const vuetify = createVuetify({
  theme: {
    defaultTheme: "myAllBlackTheme",
    themes: {
      myAllBlackTheme,
    },
  },
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi,
      fa,
    },
  },
  components,
  directives,
})

var app = createApp(App)
app.use(vuetify)
app.use(router)
app.mount('#app')

