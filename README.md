# Harvest

Parses HAR (HTTP Archive) files to extract information from them, including saving files collected during a debugging session.

The HTTP Archive format, or HAR, is a JSON-formatted archive file format for logging of a web browser’s interaction with a site.
The common extension for these files is `.har`.

## Supported Commands

### `save`

- `-include`:
      only include URLs containing this substring
- `-minbytes`:
      only include files greater than this
- `-renumber`:
      `true` to renumber; `false` to keep original filenames (default `true`)
- `-type`:
      only include matching MIME types


## Examples

- Save JSON files
  ```shell
  harvest save --type=application/json -renumber=false chimbori.com.har
  ```

- Save all images (but no other files), and rename/renumber them (to avoid meaningless filenames)
  ```shell
  harvest save --type=image/ -renumber=true chimbori.com.har
  ```

- Save only WebP images, maintaining original filenames
  ```shell
  harvest save --type=image/webp -renumber=false chimbori.com.har
  ```

- Save images from the media subdirectory
  ```shell
  harvest save --type=image/ --include=/media/ chimbori.com.har
  ```

## License

    Copyright 2024, Chimbori.

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
