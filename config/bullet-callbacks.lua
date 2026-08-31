-- Bullet callbacks -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --

local initCallbacks = {
	setVelocity = function(id, bullet, sender)
		bullet.velocity = vmath.rotate(bullet.rotation, vmath.vector3(0, bullet.speed, 0))
	end,
	setOffset = function(id, bullet, sender)
		bullet.position = bullet.position + vmath.rotate(bullet.rotation, vmath.vector3(0, bullet.offset, 0))
	end,
	recoil = function(id, bullet, sender)
		local senderVelocity = go.get(sender,"velocity")
		go.set(sender, "velocity", senderVelocity - bullet.velocity * 0.1)
	end
}

local updateCallbacks = {
	move = function(id, bullet, dt)
		bullet.position = bullet.position + bullet.velocity * dt
	end,
}

local impactCallbacks = {
	moveToEdge = function(id, bullet, bullets, event, other_event, type)
		bullet.position = bullet.position + event.normal * event.distance
	end,
	bounce = function(id, bullet, bullets, event, other_event, type)
		if bullet.bounces > 0 and (type == hash("wall") or type == hash("border")) then
			bullet.doNotDestroy = true
			bullet.velocity = bullet.velocity - 2 * vmath.dot(bullet.velocity, event.normal) * event.normal
			bullet.bounces = bullet.bounces - 1
			bullet.rotation = vmath.quat_from_to(vmath.vector3(0, 1, 0), vmath.normalize(bullet.velocity))
		end
	end,
	damageIfPlayer = function(id, bullet, bullets, event, other_event, type)
		if type == hash("player_damage") and not go.get(msg.url(nil, other_event.id, "character"), "shieldExport") then
			msg.post(other_event.id, "damage", {damage = bullet.damage})
		elseif type == hash("player_damage") then
			bullet.velocity = -bullet.velocity
			bullet.doNotDestroy = true
		end
	end
}


-- Bullet callbacks end-- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --

return {init = initCallbacks, update = updateCallbacks, impact = impactCallbacks}

