embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/gameplay/player/character.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
  scale {
    x: 0.3
    y: 0.3
    z: -0.3
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"bullet\"\n"
  "mask: \"player_damage\"\n"
  "mask: \"bullet\"\n"
  "mask: \"wall\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 2.7713976\n"
  "  data: 13.900232\n"
  "  data: 10.0\n"
  "}\n"
  "bullet: true\n"
  "event_collision: false\n"
  "event_contact: false\n"
  ""
}
