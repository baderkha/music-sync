# Music-Sync
## Author Ahmad Baderkhan

## Objectives

- A project that aims to provide an easy performant way to sync playlist from point a -> b.
- Extensible Singleton API that can access any provider
- Algorithm to sync playlist a to b will be freely available

## Backstory

There is a commericial solution called soundiz but given how slow it is and how the scheduling aspect works... i was wondering if there was a way to make this quicker and more responsive.

## Abstractions

The code is meant to be written in a consumable style. My hope is that i can create a full featured music api that can serve different providers with a unified api that allows querying songs,querying playlists, and doing basic crud operations on them

## Architecture

<center> <h3> Backend </h3> </center>

<h4> Language </h4> 
The backend will be written in go which will host the user with a restful api that has the support for copying playlists. 

<h4> Infrastructure </h4> 
The idea is to have it deployable on a FaaS kind of application that can schedule the copying job via message brokers. Which will allow for high avialability/scale.

<h4> Storage </h4>

The applicaiton should not store too much state except for caching the user's playlist syncs ...etc.

The database that will be used in a production setting will probably be mysql given its high read performance and schema flexiblity. In terms of Hosting probably something like planetscale would be perfect.


<h4> Local Testing </h4>
Locally We will ensure we use depedencies that do not need cloud involvement. For Local Testing we will mock the message brokers by using Channels to mimic the pub sub behavior and for the db we can opt for a  dockerized mysql soltuion or  a sqlite. Given how simple the models are sqlite might be the better solution here.


<center> <h3> Front End </h3> </center>