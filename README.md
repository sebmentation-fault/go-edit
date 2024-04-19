# go-edit, Go Text Editor

My first golang project -- let's see if I am able to learn it well.

## Requirements

* terminal user interface
* open 1 file of choice (maybe more later)
* file can be saved

## Wireframe

Edit Mode

```

  Terminal Window
  +--------------------------------------------------+-------+
  | File Name                                        | [ESC] |
  +--------------------------------------------------+-------+
  |                                                          |
  |   Lorem ipsum dolor sit amet, officia excepteur ex       |
  | fugiat reprehenderit enim labore culpa sint ad nisi.     |
  |                                                          |
  |   Lorem pariatur mollit ex esse exercitation amet.       |
  | Nisi anim cupidatat excepteur officia.                   |
  |                                                          |
  |   Reprehenderit nostrud nostrud ipsum Lorem est aliquip  |
  | amet voluptate voluptate dolor minim nulla est           |
  | proident. Nostrud officia pariatur ut officia.           |
  |                                                          |
  |   Sit irure elit esse ea nulla sunt ex occaecat          |
  | reprehenderit commodo officia dolor Lorem duis laboris   |
  | cupidatat officia voluptate.                             |
  |                                                          |
  |   Culpa proident adipisicing id nulla nisi laboris ex in |
  | Lorem sunt duis officia eiusmod.                         |
  |                                                          |
  +----------------------------------------------------------+

```

Escaped Mode

```

  Terminal Window
  +----------------------------------------+--------+--------+
  | File Name                              | [Q]iut | [S]ave |
  +----------------------------------------+--------+--------+
  |                                                          |
  |   Lorem ipsum dolor sit amet, officia excepteur ex       |
  | fugiat reprehenderit enim labore culpa sint ad nisi.     |
  |                                                          |
  |   Lorem pariatur mollit ex esse exercitation amet.       |
  | Nisi anim cupidatat excepteur officia.                   |
  |                                                          |
  |   Reprehenderit nostrud nostrud ipsum Lorem est aliquip  |
  | amet voluptate voluptate dolor minim nulla est           |
  | proident. Nostrud officia pariatur ut officia.           |
  |                                                          |
  |   Sit irure elit esse ea nulla sunt ex occaecat          |
  | reprehenderit commodo officia dolor Lorem duis laboris   |
  | cupidatat officia voluptate.                             |
  |                                                          |
  |   Culpa proident adipisicing id nulla nisi laboris ex in |
  | Lorem sunt duis officia eiusmod.                         |
  |                                                          |
  +----------------------------------------------------------+

```

## Usage and Specifications

* user opens file `goted <filename>`
* edit mode:
  * user enters editing mode by default -- start typing away
  * if user presses <esc> , they are in Escaped mode
* escaped mode:
  * <q> - quit the application (if not saved since change made, prompt the user if they are sure)
  * <s> - save the file
  * <enter> - enter edit mode again
* terminal windows that are less than 25 characters wide are not supported -- the file name, quit 
  and save banner won't have space!

> Yes, I'm a fan of Vim. Maybe, with time, escaped mode becomes visual mode and I make some 
  commands...


