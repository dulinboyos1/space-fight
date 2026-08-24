embedded_components {
  id: "background"
  type: "sprite"
  data: "default_animation: \"Modifier Background\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/gameplay/modifiers/modifiers.atlas\"\n"
  "}\n"
  ""
  position {
    z: -0.1
  }
}
embedded_components {
  id: "title"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"Modifier Title\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    y: 518.0
  }
  scale {
    x: 5.547548
    y: 5.547548
    z: 5.547548
  }
}
embedded_components {
  id: "subtitle"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"Subtitle\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    y: 419.0
  }
  scale {
    x: 3.5
    y: 3.5
    z: 3.5
  }
}
embedded_components {
  id: "desc"
  type: "label"
  data: "size {\n"
  "  x: 250.0\n"
  "  y: 100.0\n"
  "}\n"
  "pivot: PIVOT_N\n"
  "line_break: true\n"
  "text: \"Interesting modifier description This is a very interesting description\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    y: 13.0
  }
  scale {
    x: 3.0
    y: 3.0
    z: 3.0
  }
}
embedded_components {
  id: "img"
  type: "sprite"
  data: "default_animation: \"example\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/gameplay/modifiers/modifier-thumbnails.atlas\"\n"
  "}\n"
  ""
  position {
    y: 194.0
    z: 0.1
  }
  scale {
    x: 3.198272
    y: 3.198272
    z: 3.198272
  }
}
embedded_components {
  id: "properties"
  type: "label"
  data: "size {\n"
  "  x: 250.0\n"
  "  y: 32.0\n"
  "}\n"
  "pivot: PIVOT_N\n"
  "text: \"example: +10%\\n"
  "\"\n"
  "  \"health: -100\\n"
  "\"\n"
  "  \"\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    y: -140.0
  }
  scale {
    x: 4.0
    y: 4.0
    z: 4.0
  }
}
embedded_components {
  id: "menu-border"
  type: "sprite"
  data: "default_animation: \"modifier-border\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 835.0\n"
  "  y: 1220.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/menu-navigator/menu-navigator.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "cover"
  type: "sprite"
  data: "default_animation: \"Modifier Cover\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/gameplay/modifiers/modifiers.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
}
