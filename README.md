# rss-transform

## Status

Mostly incomplete, WIP. I've got a simple framework that does nothing, and I'm
working on defining the contracts between the individual components.

## Summary

An ETL pipeline for consuming, transforming and republishing RSS feeds.

## Overview

<dl>
  <dt>

[RSS](https://en.wikipedia.org/wiki/RSS)

  </dt>
  <dd>An XML format for publishing information on the internet. An RSS feed is
  an array of Items, ordered most-recent-first.</dd>
  <dt>

[ETL](https://en.wikipedia.org/wiki/Extract,_transform,_load)

  </dt>
  <dd>A common pattern for migrating data between systems, consisting of
  Extract (get data from a source), Transform (transform that data into a
  format your destination system can accept) and Load (publish that data to the
  destination system).</dd>
  </dl>

For the `rss-transform` project, we define the ETL stages like this:

* *Extract*: Collect data from a source system. This could be an RSS feed,
  a database, a file or an API. But the data should be returned as discrete
  Items.
* *Transform*: Perform transformations on the extracted Items to prepare them
  for publishing as an RSS feed. Common steps are cleaning (removing
  duplicates, correcting errors), decorating (adding extra information),
  normalizing (converting to a common format or structure), filtering,
  splitting, and merging.
* *Load*: Our destination is normally just an XML file, but it's possible to
  write plugins for any destination system.

The `rss-transform` project consists of a command-line tool (`rss-pipeline`)
that manages the pipelines and Items, and a collection of plugins that perform
individual ETL operations. Plugins are standalone programs that accept JSON on
stdin and return JSON on stdout. This makes creating new plugins in any
language as easy as possible - they can even be shell scripts.

## Configuring `rss-pipeline`

The tool takes one command line argument, the location of its config directory:

    `rss-pipeline --config /etc/rss-pipeline`

It looks in this directory for a file named `config.yml`. All other `*.yml`
files in the directory are assumed to be pipeline definitions, and loaded. Any
other file in the directory is ignored. It does not recurse into child directories.

```yaml
MaxAge: 30
MaxNumItems: 1000
```

<dl>
  <dt>MaxAge</dt>
  <dd>If individual items are older than 30 days, delete them.</dd>
  <dt>MaxNumItems</dt>
  <dd>If we have more than 1000 items, delete the oldest.</dd>
</dl>

If both these parameters are set, items that meet either criteria are deleted.

If neither of these parameters are set, nothing is ever deleted. There are
currently no tools for pruning these items manually.

## Configuring a pipeline

Each pipeline is defined in its own YAML file. Here's the simplest possible
pipeline:

```yaml
Name: Echo
Description: Republishes an RSS feed.
SleepDuration: 1800
Pipeline:
  Extract:
    Exec: "./get-rss-feed"
    Config:
      Url: "https://feeds.bbci.co.uk/news/world/rss.xml"
  Transform:
  Load:
    -
      Exec: "./publish-rss-feed"
      NumItems: 100
      Config:
        Filename: "./output/bbc-world-news.xml"
```

This defines a pipeline called "Echo" that runs every 30 minutes
(1800 seconds), gets the BBC World RSS feed, does nothing to
it, and saves it to a file locally.

The important bit's the `Pipeline` key. This is what's going on under the hood:

1. Every `SleepDuration` seconds, `rss-pipeline` runs the program `./get-rss-feed` and passes
   it the following JSON on stdin:

   ```json
   [
     {
       "Config": {
         "Url": "https://feeds.bbci.co.uk/news/world/rss.xml"
       }
     }
   ]
   ```
2. `./get-rss-feed` hits the URL, grabs the document, and returns its contents as an array of Items
   on stdout:

   ```json
   [
     {
       "Result": "@todo"
     },
     {
       "Data": [
         {
           "@todo": "@todo"
         },
         {
           "@todo": "@todo"
         }
       ]
     }
   ]
   ```
3. If there are NO new items since last time, the pipeline stops processing and
   goes back to sleep for 1800 seconds.
4. If there ARE new items, nothing is done to them (because the `Transform`
   step is empty) and up to `NumItems` of them are passed to `./PublishRssFeed`
   as JSON, most recent first. `rss-pipeline` stores Items once processed, so
   we can publish more than were in the original feed.

   ```json
   [
     {
       "Config": {
         "Filename": "./output/bbc-world-news.xml"
       }
     },
     {
       "Data": [
         {
           "@todo": "@todo"
         },
         {
           "@todo": "@todo"
         }
       ]
     }
   ]
   ```
5. `./PublishRssFeed` creates or overwrites `Filename` with the data formatted
   as an RSS feed, and returns success or failure:
   ```json
   [
     {
       "Result": "@todo"
     }
   ]
   ```
6. The pipeline goes to sleep for 1800 seconds, then the whole process starts over.

Notes:

* It's not possible to have more than one `Extract` plugin. If you want to
  decorate Items with additional data, use a `Transform` plugin
* It's expected to have more than one `Process` plugin
* It's possible to have more than one `Load` plugin
* `rss-pipeline` doesn't understand the parameters it's passing to the plugins
  via the `Config` key, it just does it
* `Extract` plugins don't get passed `Data` keys
* `Load` plugins don't return `Data` keys
* All plugins return `Result` keys
* `Transform` plugins are passed individual Items, not the whole array. This is
  because `rss-pipeline` keeps track of all the Items so each one only has to
  be processed once.
* If `NumItems` is 0, no items are sent.
* If `NumItems` is -1, all available items are sent.

## Inspiration

[RSS-Bridge](https://github.com/RSS-Bridge/rss-bridge) and the [Drupal Migrate API](https://www.drupal.org/docs/drupal-apis/migrate-api)
