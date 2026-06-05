embedded_components {
  id: "remove-button"
  type: "sprite"
  data: "default_animation: \"Remove Button\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/start-menu/start-menu.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1301.0
  }
}
embedded_components {
  id: "label"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "pivot: PIVOT_W\n"
  "text: \" = Player [_]\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  scale {
    x: 5.5
    y: 5.5
    z: 5.5
  }
}
embedded_components {
  id: "menu-border"
  type: "sprite"
  data: "default_animation: \"Remove Button Highlighter\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/menu-navigator/menu-navigator.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1302.0
    y: 1.0
  }
}
