<template>
  <v-container>
    <v-row class="text-center">
      <div>
        <p v-if="login">You are already logged in.</p>
        <p v-else>Not logged in</p>
      </div>
      <div>
        ID: {{ this.id }} Username: {{ this.username }}
      </div>
      <div id="app">
        {{ info }}
      </div>
    </v-row>
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
      <v-btn
        color="primary"
        dark
        v-on:click="getTimelines"
      >
      Update
      </v-btn>
    </v-row>
    <v-row>
    <v-simple-table>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">
            Author
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
          v-for="item in tweets"
          :key="item.name"
        >
          <td>{{ item.author_id }}</td>
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
        <!-- Todo: Rules -->
        <v-text-field
          v-model="tweet_text"
          :counter="140"
          label="Your Tweet"
          required
        ></v-text-field>
      </v-col>
      <v-col>
      <v-btn
        color="primary"
        dark
        v-on:click="postTweet"
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
      <template>
        <v-btn
          color="primary"
          dark
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
  import axios from 'axios';

  export default {
    name: 'Top',

    data () {
      return {
        // 定数
        APISV: "http://localhost:8081",
        ENDPOINT_TIMELINES: "/timelines",
        ENDPOINT_TWEET:     "/tweet",
        ENDPOINT_SEARCH:    "/search",

        dialog: false,
        info: null,
        counter: 0,
        id: 0,
        username: "",
        tweet_text: "",
        login: false,
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
        ],
        tweets: [],
      }
    },
    mounted () {
      console.log('mounted');
      console.log(this.login);
      if (this.$route.query.id && this.login == false) {
        this.login = true
        this.id = this.$route.query.id
        this.username = this.$route.query.username
      }
    },
    methods: {
      getTimelines: function () {
        axios
        .get(this.APISV+this.ENDPOINT_TIMELINES, {
            params: {
              id: this.id
            }
          })
          .then(response => {
            this.tweets = []
            for(var i=0;i<response.data.data.length;i++){
              console.log(response.data.data[i])
              this.tweets.push({"author_id":"","text":""})
              this.tweets[i].author_id = response.data.data[i].author_id
              this.tweets[i].text = response.data.data[i].text
            }
            //if (this.tweets.length == 0) {
            //}
            //this.tweets.push({"name":"aaa","text":"April"})
          })
      },
      postTweet: function () {
        //console.log(this.tweet_text);
        axios
          .post(this.APISV+this.ENDPOINT_TWEET, {
            text: this.tweet_text
          })
          .then(response => {
            console.log(response.data)
          })
      },
      searchTweets: function () {
        axios
          .post(this.APISV+this.ENDPOINT_SEARCH)
          .then(response => {
            this.tweets[0].author_id = response.data.data[0].author_id
          })
      },
    },
  }
</script>
