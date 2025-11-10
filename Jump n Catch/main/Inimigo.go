components {
  id: "Inimigo.script"
  component: "/main/assets/scripts/Inimigo.script"
}
embedded_components {
  id: "Inimigo"
  type: "sprite"
  data: "default_animation: \"inimigo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tiles/Player.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.506186
    y: 0.619579
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"Inimigo\"\n"
  "mask: \"ground\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 5.3809876\n"
  "  data: 6.40611\n"
  "  data: 14.591551\n"
  "}\n"
  ""
}
