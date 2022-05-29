<template>
  <v-container>
    <v-row class="text-center">
      <v-toolbar
        class="mb-2"
        color="primary"
        dark
        flat
      >
        <v-toolbar-title>
          Timelines
        </v-toolbar-title>
      </v-toolbar>
    </v-row>
    <v-row>
    <v-simple-table>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            Tweet
          </th>
          <th class="text-left">
            Like
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in desserts"
          :key="item.name"
        >
          <td>{{ item.name }}</td>
          <td>{{ item.text }}</td>
          <td>heart</td>
        </tr>
      </tbody>
    </template>
    </v-simple-table>
    </v-row>

    <v-row class="text-center">
      <v-toolbar
        class="mb-2"
        color="primary"
        dark
        flat
      >
        <v-toolbar-title>
          Tweet
        </v-toolbar-title>
      </v-toolbar>
    </v-row>

    <v-row>
      <v-col>
        <v-text-field
          v-model="tweet"
          :rules="tweetRules"
          :counter="140"
          label="Your Tweet"
          required
        ></v-text-field>
      </v-col>
      <v-col>
      <v-btn
        color="primary"
        dark
        v-bind="attrs"
        v-on="on"
      >
        Tweet!
      </v-btn>
      </v-col>
    </v-row>

    <v-row>
      <v-toolbar
        class="mb-2"
        color="primary"
        dark
        flat
      >
        <v-toolbar-title>
          Search Tweets
        </v-toolbar-title>
      </v-toolbar>
    </v-row>
    <v-row>
      <v-dialog
        v-model="dialog"
        width="500"
      >
        <template v-slot:activator="{on, attrs }">
        <v-btn
          color="primary"
          dark
          v-bind="attrs"
          v-on="on"
        >
          Click Me
        </v-btn>
        </template>
        <v-card>
        <v-card-title class="text-h5 grey lighten-2">
          Privacy Policy
        </v-card-title>

        <v-card-text>
          Lorem .
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            text
            @click="dialog = false"
          >
            I accept
          </v-btn>
        </v-card-actions>
      </v-card>
      </v-dialog>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'HelloWorld',

    data () {
      return {
        dialog: false,
        desserts: [
          {
            name: 'Frozen Yogurt',
            text: '159',
          },
                    {
            name: 'Ice cream sandwich',
            text: '237',
          },
          {
            name: 'Eclair',
            text: '262',
          },
        ]
      }
    },
  }
</script>
