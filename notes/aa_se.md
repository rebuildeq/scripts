id|name|desc
--|--|--
0|SE_CurrentHP|implemented - Heals and nukes, repeates every tic if in a buff
1|SE_ArmorClass|implemented
2|SE_ATK|implemented
3|SE_MovementSpeed|implemented - SoW, SoC, etc
4|SE_STR|implemented
5|SE_DEX|implemented
6|SE_AGI|implemented
7|SE_STA|implemented
8|SE_INT|implemented
9|SE_WIS|implemented
10|SE_CHA|implemented - used as a spacer
11|SE_AttackSpeed|implemented
12|SE_Invisibility|implemented - TO DO: Implemented Invisiblity Levels
13|SE_SeeInvis|implemented - TO DO: Implemented See Invisiblity Levels
14|SE_WaterBreathing|implemented
15|SE_CurrentMana|implemented
16|SE_NPCFrenzy|not used
17|SE_NPCAwareness|not used
18|SE_Lull|implemented - Reaction Radius
19|SE_AddFaction|implemented - Alliance line
20|SE_Blind|implemented
21|SE_Stun|implemented
22|SE_Charm|implemented
23|SE_Fear|implemented
24|SE_Stamina|implemented - Invigor and such
25|SE_BindAffinity|implemented - TO DO: Implement 2nd and 3rd Recall (value 2,3 ect). Sets additional bind points.
26|SE_Gate|implemented - Gate to bind point
27|SE_CancelMagic|implemented
28|SE_InvisVsUndead|implemented
29|SE_InvisVsAnimals|implemented
30|SE_ChangeFrenzyRad|implemented - Pacify
31|SE_Mez|implemented
32|SE_SummonItem|implemented
33|SE_SummonPet|implemented
34|SE_Confuse|not used (Nimbus of Temporal Rifting) ?
35|SE_DiseaseCounter|implemented
36|SE_PoisonCounter|implemented
37|SE_DetectHostile|not used
38|SE_DetectMagic|not used
39|SE_TwinCastBlocker|implemented - If present in spell, then the spell can not be twincast.
40|SE_DivineAura|implemented
41|SE_Destroy|implemented - Disintegrate, Banishment of Shadows
42|SE_ShadowStep|implemented
43|SE_Berserk|implemented (*not used in any known live spell) Makes client 'Berserk' giving crip blow chance.
44|SE_Lycanthropy|implemented
45|SE_Vampirism|implemented (*not used in any known live spell) Stackable lifetap from melee.
46|SE_ResistFire|implemented
47|SE_ResistCold|implemented
48|SE_ResistPoison|implemented
49|SE_ResistDisease|implemented
50|SE_ResistMagic|implemented
51|SE_DetectTraps|not used
52|SE_SenseDead|implemented
53|SE_SenseSummoned|implemented
54|SE_SenseAnimals|implemented
55|SE_Rune|implemented
56|SE_TrueNorth|implemented
57|SE_Levitate|implemented
58|SE_Illusion|implemented
59|SE_DamageShield|implemented
60|SE_TransferItem|not used
61|SE_Identify|implemented
62|SE_ItemID|not used
63|SE_WipeHateList|implemented, @Memblur, chance to wipe hate list of target, base: pct chance, limit: none, max: ? (not implemented), Note: caster level and CHA add to pct chance
64|SE_SpinTarget|implemented - TO DO: Not sure stun portion is working correctly
65|SE_InfraVision|implemented
66|SE_UltraVision|implemented
67|SE_EyeOfZomm|implemented
68|SE_ReclaimPet|implemented
69|SE_TotalHP|implemented
70|SE_CorpseBomb|not used
71|SE_NecPet|implemented
72|SE_PreserveCorpse|not used
73|SE_BindSight|implemented, @Vision, see through the eyes of your target, click off buff to end effect, base: 1, limit: none, max: none
74|SE_FeignDeath|implemented
75|SE_VoiceGraft|implemented
76|SE_Sentinel|*not implemented?(just seems to send a message)
77|SE_LocateCorpse|implemented
78|SE_AbsorbMagicAtt|implemented - Rune for spells
79|SE_CurrentHPOnce|implemented - Heals and nukes, non-repeating if in a buff
80|SE_EnchantLight|not used
81|SE_Revive|implemented - Resurrect
82|SE_SummonPC|implemented
83|SE_Teleport|implemented
84|SE_TossUp|implemented - Gravity Flux
85|SE_WeaponProc|implemented - i.e. Call of Fire
86|SE_Harmony|implemented
87|SE_MagnifyVision|implemented - Telescope
88|SE_Succor|implemented - Evacuate/Succor lines
89|SE_ModelSize|implemented - Shrink, Growth
90|SE_Cloak|*not implemented - Used in only 2 spells
91|SE_SummonCorpse|implemented
92|SE_InstantHate|implemented - add hate
93|SE_StopRain|implemented - Wake of Karana
94|SE_NegateIfCombat|implemented
95|SE_Sacrifice|implemented
96|SE_Silence|implemented
97|SE_ManaPool|implemented
98|SE_AttackSpeed2|implemented - Melody of Ervaj
99|SE_Root|implemented
100|SE_HealOverTime|implemented
101|SE_CompleteHeal|implemented
102|SE_Fearless|implemented - Valiant Companion
103|SE_CallPet|implemented - Summon Companion
104|SE_Translocate|implemented
105|SE_AntiGate|implemented - Translocational Anchor
106|SE_SummonBSTPet|implemented
107|SE_AlterNPCLevel|implemented - not used on live
108|SE_Familiar|implemented
109|SE_SummonItemIntoBag|implemented - summons stuff into container
110|SE_IncreaseArchery|not used
111|SE_ResistAll|implemented - Note: Physical Resists are not modified by this effect.
112|SE_CastingLevel|implemented
#define	SE_SummonHorse					113	implemented
114|SE_ChangeAggro|implemented - Hate modifing buffs(ie horrifying visage)
115|SE_Hunger|implemented - Song of Sustenance
116|SE_CurseCounter|implemented
117|SE_MagicWeapon|implemented - makes weapon magical
118|SE_Amplification|implemented, @Song, stackable singing mod, base: mod%, limit: none, max: none, Note: Can focus itself.
119|SE_AttackSpeed3|implemented
120|SE_HealRate|implemented - reduces healing by a %
121|SE_ReverseDS|implemented
122|SE_ReduceSkill|not implemented    TODO: Now used on live, decreases skills by percent
123|SE_Screech|implemented Spell Blocker(If have buff with value +1 will block any effect with -1)
124|SE_ImprovedDamage|implemented
125|SE_ImprovedHeal|implemented
126|SE_SpellResistReduction|implemented
127|SE_IncreaseSpellHaste|implemented, @Fc, On Caster, cast time mod pct, base: pct
128|SE_IncreaseSpellDuration|implemented, @Fc, On Caster, spell duration mod pct, base: pct
129|SE_IncreaseRange|implemented, @Fc, On Caster, spell range mod pct, base: pct
130|SE_SpellHateMod|implemented, @Fc, On Caster, spell hate mod pct, base: min pct, limit: max pct
131|SE_ReduceReagentCost|implemented, @Fc, On Caster, do not consume reagent pct chance, base: min pct, limit: max pct
132|SE_ReduceManaCost|implemented, @Fc, On Caster, reduce mana cost by pct, base: min pct, limt: max pct
133|SE_FcStunTimeMod|implemented, @Fc, On Caster, spell range mod pct, base: pct
134|SE_LimitMaxLevel|implemented, @Ff, Max level of spell that can be focused, if base2 then decrease effectiviness by base2 % per level over max, base:  lv, base2: effectiveness pct
135|SE_LimitResist|implemented, @Ff, Resist Type(s) that a spell focus can require or exclude, base1: resist type, Include: Positive Exclude: Negative
136|SE_LimitTarget|implemented, @Ff, Target Type(s) that a spell focus can require or exclude, base1: target type, Include: Positive Exclude: Negative
137|SE_LimitEffect|implemented, @Ff, Spell effect(s) that a spell focus can require or exclude, base1: SPA id, Include: Positive Exclude: Negative
138|SE_LimitSpellType|implemented, @Ff, Only allow focus spells that are Beneficial or Detrimental, base1: 0=det 1=bene
139|SE_LimitSpell|implemented, @Ff, Specific spell id(s) that a spell focus can require or exclude, base1: SPA id, Include: Positive Exclude: Negative
140|SE_LimitMinDur|implemented, @Ff, Mininum duration of spell that can be focused, base1: tics
141|SE_LimitInstant|implemented, @Ff, Include or exclude if an isntant cast spell can be focused, base1: 0=Exclude if Instant 1=Allow only if Instant
142|SE_LimitMinLevel|implemented, @Ff, Mininum level of spell that can be focused, base1: lv
143|SE_LimitCastTimeMin|implemented, @Ff, Mininum cast time of spell that can be focused, base1: milliseconds
144|SE_LimitCastTimeMax|implemented, @Ff, Max cast time of spell that can be focused, base1: milliseconds
145|SE_Teleport2|implemented - Banishment of the Pantheon
146|SE_ElectricityResist|*not implemented TODO: Now used on live, xyz for teleport spells? also in temp pets?
147|SE_PercentalHeal|implemented
148|SE_StackingCommand_Block|implemented?
149|SE_StackingCommand_Overwrite|implemented?
150|SE_DeathSave|implemented
151|SE_SuspendPet|implemented, @Pet, allow caster to have an extra suspended pet, base: 0=no buffs/items 1=buffs+items, limit: none, max: none
152|SE_TemporaryPets|implemented
153|SE_BalanceHP|implemented
154|SE_DispelDetrimental|implemented, @Dispel, removes only detrimental effects on a target, base: pct chance (950=95%), limit: none, max: none
155|SE_SpellCritDmgIncrease|implemented - no known live spells use this currently
156|SE_IllusionCopy|implemented - Deception
157|SE_SpellDamageShield|implemented, @DS, causes non-melee damage on caster of a spell, base: Amt DS (negative), limit: none, max: unknown (same as base but +)
158|SE_Reflect|implemented, @SpellMisc, reflect casted detrimental spell back at caster, base: chance pct, limit: resist modifier (positive value reduces resists), max: pct of base dmg mod (50=50pct of base)
159|SE_AllStats|implemented
160|SE_MakeDrunk|*not implemented - Effect works entirely client side (Should check against tolerance)
161|SE_MitigateSpellDamage|implemented, @Runes, mitigate incoming spell damage by percentage until rune fades, base: percent mitigation, limit: max dmg absorbed per hit, max: rune amt, Note: If placed on item or AA, will provide stackable percent mitigation.
162|SE_MitigateMeleeDamage|implemented - rune with max value
163|SE_NegateAttacks|implemented
164|SE_AppraiseLDonChest|implemented
165|SE_DisarmLDoNTrap|implemented
166|SE_UnlockLDoNChest|implemented
167|SE_PetPowerIncrease|implemented, @Fc, On Caster, pet power mod, base: value
168|SE_MeleeMitigation|implemented
169|SE_CriticalHitChance|implemented
170|SE_SpellCritChance|implemented
171|SE_CrippBlowChance|implemented
172|SE_AvoidMeleeChance|implemented
173|SE_RiposteChance|implemented
174|SE_DodgeChance|implemented
175|SE_ParryChance|implemented
176|SE_DualWieldChance|implemented
177|SE_DoubleAttackChance|implemented
178|SE_MeleeLifetap|implemented
179|SE_AllInstrumentMod|implemented, @Song, set mod for ALL instrument/singing skills that will be used if higher then item mods, base: mod%, limit: none, max: none
180|SE_ResistSpellChance|implemented
181|SE_ResistFearChance|implemented
182|SE_HundredHands|implemented
183|SE_MeleeSkillCheck|implemented
184|SE_HitChance|implemented
185|SE_DamageModifier|implemented
186|SE_MinDamageModifier|implemented
187|SE_BalanceMana|implemented - Balances party mana
188|SE_IncreaseBlockChance|implemented
189|SE_CurrentEndurance|implemented
190|SE_EndurancePool|implemented
191|SE_Amnesia|implemented - Silence vs Melee Effect
192|SE_Hate|implemented - Instant and hate over time.
193|SE_SkillAttack|implemented,
194|SE_FadingMemories|implemented, @Aggro, Remove from hate lists and make invisible. Can set max level of NPCs that can be affected. base: success chance, limit: max level (ROF2), max: max level (modern client), Note: Support for max level requires Rule (Spells, UseFadingMemoriesMaxLevel) to be true. If used from limit field, then it set as the level, ie. max level of 75 would use limit value of 75. If set from max field, max level 75 would use max value of 1075, if you want to set it so it checks a level range above the spell target then for it to only work on mobs 5 levels or below you set max value to 5.
195|SE_StunResist|implemented
196|SE_StrikeThrough|implemented
197|SE_SkillDamageTaken|implemented
198|SE_CurrentEnduranceOnce|implemented
199|SE_Taunt|implemented - % chance to taunt the target
200|SE_ProcChance|implemented
201|SE_RangedProc|implemented
202|SE_IllusionOther|implemented - Project Illusion
203|SE_MassGroupBuff|implemented
204|SE_GroupFearImmunity|implemented - (Does not use bonus)
205|SE_Rampage|implemented, @Combat Instant, Perform a primary slot combat rounds on all creatures within a 40 foot radius, base: number of attack rounds, limit: max entities hit per round, max: none, Note: AE range is 40 by default. Custom: Set field 'aoe_range' to override default. Adding additional attacks and hit count limit.
206|SE_AETaunt|implemented
207|SE_FleshToBone|implemented
208|SE_PurgePoison|not used
209|SE_DispelBeneficial|implemented, @Dispel, removes only beneficial effects on a target, base: pct chance (950=95%), limit: none, max: none
210|SE_PetShield|implmented, @ShieldAbility, allows pet to 'shield' owner for 50 pct of damage taken for a duration, base: Time multiplier 1=12 seconds, 2=24 ect, limit: mitigation on pet owner override (not on live), max: mitigation on pet overide (not on live)
211|SE_AEMelee|implemented TO DO: Implement to allow NPC use (client only atm).
212|SE_FrenziedDevastation|implemented - increase spell criticals + all DD spells cast 2x mana.
213|SE_PetMaxHP|implemented[AA] - increases the maximum hit points of your pet
214|SE_MaxHPChange|implemented
215|SE_PetAvoidance|implemented[AA] - increases pet ability to avoid melee damage
216|SE_Accuracy|implemented
217|SE_HeadShot|implemented - ability to head shot (base2 = damage)
218|SE_PetCriticalHit|implemented[AA] - gives pets a baseline critical hit chance
219|SE_SlayUndead|implemented - Allow extra damage against undead (base1 = rate, base2 = damage mod).
220|SE_SkillDamageAmount|implemented
221|SE_Packrat|implemented as bonus
222|SE_BlockBehind|implemented - Chance to block from behind (with our without Shield)
223|SE_DoubleRiposte|implemented - Chance to double riposte [not used on live]
#define	SE_GiveDoubleRiposte			224 implemented[AA]
225|SE_GiveDoubleAttack|implemented[AA] - Allow any class to double attack with set chance.
226|SE_TwoHandBash|*not implemented as bonus
227|SE_ReduceSkillTimer|implemented
228|SE_ReduceFallDamage|implented - reduce the damage that you take from falling
229|SE_PersistantCasting|implemented
230|SE_ExtendedShielding|implemented, @ShieldAbility, extends the range of your /shield ability by an amount of distance, base: distance units, limit: none, max: none
231|SE_StunBashChance|implemented - increase chance to stun from bash.
232|SE_DivineSave|implemented (base1 == % chance on death to insta-res) (base2 == spell cast on save)
233|SE_Metabolism|implemented - Modifies food/drink consumption rates.
234|SE_ReduceApplyPoisonTime|not implemented as bonus - reduces the time to apply poison
#define	SE_ChannelChanceSpells			235 implemented[AA] - chance to channel from SPELLS *No longer used on live.
236|SE_FreePet|not used
237|SE_GivePetGroupTarget|implemented[AA] - (Pet Affinity)
238|SE_IllusionPersistence|implemented - lends persistence to your illusionary disguises, causing them to last until you die or the illusion is forcibly removed.
239|SE_FeignedCastOnChance|implemented - ability gives you an increasing chance for your feigned deaths to not be revealed by spells cast upon you.
240|SE_StringUnbreakable|not used [Likely related to above - you become immune to feign breaking on a resisted spell and have a good chance of feigning through a spell that successfully lands upon you.]
241|SE_ImprovedReclaimEnergy|implemented - increase the amount of mana returned to you when reclaiming your pet.
242|SE_IncreaseChanceMemwipe|implemented - @Memblur, increases the chance to wipe hate with memory blurr, base: chance pct, limit: none, max: none, Note: Mods final blur chance after other bonuses added.
243|SE_CharmBreakChance|implemented - Total Domination
#define	SE_RootBreakChance				244	implemented[AA] reduce the chance that your root will break.
245|SE_TrapCircumvention|implemented, @Traps, decreases the chance that you will set off a trap when opening a chest or other similar container by percentage, base: chance modifer, limit: none, max: none
246|SE_SetBreathLevel|*not implemented as bonus
247|SE_RaiseSkillCap|implemented[AA] - adds skill over the skill cap.
248|SE_SecondaryForte|not implemented as bonus(gives you a 2nd specialize skill that can go past 50 to 100)
249|SE_SecondaryDmgInc|implemented[AA] Allows off hand weapon to recieve a damage bonus (Sinister Strikes)
250|SE_SpellProcChance|implemented - Increase chance to proc from melee proc spells (ie Spirit of Panther)
251|SE_ConsumeProjectile|implemented[AA] - chance to not consume an arrow (ConsumeProjectile = 100)
252|SE_FrontalBackstabChance|implemented[AA] - chance to perform a full damage backstab from front.
253|SE_FrontalBackstabMinDmg|implemented[AA] - allow a frontal backstab for mininum damage.
254|SE_Blank|implemented
255|SE_ShieldDuration|implemented, , @ShieldAbility, extends the duration of your /shield ability, base: seconds, limit: none, max: none
256|SE_ShroudofStealth|implemented
257|SE_PetDiscipline|not implemented as bonus - /pet hold - official name is GivePetHold
258|SE_TripleBackstab|implemented[AA] - chance to perform a triple backstab
259|SE_CombatStability|implemented[AA] - damage mitigation
260|SE_AddSingingMod|implemented, @Song, set mod for specific instrument/singing skills that will be used if higher then item mods, base: mod%, limit: ItemType ID, max: none
261|SE_SongModCap|implemented, @Song, raise max song modifier cap, base: amt, limit: none, max: none, Note: No longer used on live
262|SE_RaiseStatCap|implemented
263|SE_TradeSkillMastery|implemented - lets you raise more than one tradeskill above master.
264|SE_HastenedAASkill|implemented
265|SE_MasteryofPast|implemented[AA] - Spells less than effect values level can not be fizzled
266|SE_ExtraAttackChance|implemented, @OffBonus, gives your double attacks a percent chance to perform an extra attack with 2-handed primary weapon, base: chance, limit: amt attacks, max: none
267|SE_AddPetCommand|implemented - sets command base2 to base1
268|SE_ReduceTradeskillFail|implemented - reduces chance to fail with given tradeskill by a percent chance
269|SE_MaxBindWound|implemented[AA] - Increase max HP you can bind wound.
270|SE_BardSongRange|implemented, @Song, increase range of beneficial bard songs, base: mod%, limit: none, max: none , Note: example Sionachie's Crescendo
271|SE_BaseMovementSpeed|implemented[AA] - mods basemove speed, doesn't stack with other move mods
272|SE_CastingLevel2|implemented
273|SE_CriticalDoTChance|implemented
274|SE_CriticalHealChance|implemented
275|SE_CriticalMend|implemented[AA] - chance to critical monk mend
276|SE_Ambidexterity|implemented[AA] - increase chance to duel weild by adding bonus 'skill'
277|SE_UnfailingDivinity|implemented[AA] - ability grants your Death Pact-type spells a second chance to successfully heal their target, also can cause said spells to do a portion of their healing value even on a complete failure.
#define	SE_FinishingBlow				278 implemented[AA] - chance to do massive damage under 10% HP (base1 = chance, base2 = damage)
279|SE_Flurry|implemented
280|SE_PetFlurry|implemented[AA]
281|SE_FeignedMinion|implemented, ability allows you to instruct your pet to feign death via the '/pet feign' command, base: succeed chance, limit: none, max: none, Note: Only implemented as an AA.
282|SE_ImprovedBindWound|implemented[AA] - increase bind wound amount by percent.
283|SE_DoubleSpecialAttack|implemented[AA] - Chance to perform second special attack as monk
284|SE_LoHSetHeal|not used
285|SE_NimbleEvasion|*not implemented - base1 = 100 for max
286|SE_FcDamageAmt|implemented, @Fc, On Caster, spell damage mod flat amt, base: amt
287|SE_SpellDurationIncByTic|implemented, @Fc, SPA: 287, SE_SpellDurationIncByTic,			On Caster, spell buff duration mod, base: tics
288|SE_SkillAttackProc|implemented, @Procs, chance to cast a spell when using a skill, base: chance, limit: skill, max: spellid, note: if used in AA the spell id is set in aa_ranks spell field, chance is calculated as 100% = value 1000.
289|SE_CastOnFadeEffect|implemented - Triggers only if fades after natural duration.
290|SE_IncreaseRunSpeedCap|implemented[AA] - increases run speed over the hard cap
291|SE_Purify|implemented, @Dispel, remove up specified amount of detiremental spells, base: amt removed, limit: none, max: none, Note: excluding charm, fear, resurrection, and revival sickness
292|SE_StrikeThrough2|implemented[AA] - increasing chance of bypassing an opponent's special defenses, such as dodge, block, parry, and riposte.
293|SE_FrontalStunResist|implemented[AA] - Reduce chance to be stunned from front. -- live descriptions sounds like this isn't limited to frontal anymore
294|SE_CriticalSpellChance|implemented - increase chance to critical hit and critical damage modifier.
295|SE_ReduceTimerSpecial|not used
296|SE_FcSpellVulnerability|implemented, @Fc, On Target, spell damage taken mod pct, base: min pct, limit: max pct
297|SE_FcDamageAmtIncoming|implemetned, @Fc, On Target, damage taken flat amt, base: amt
298|SE_ChangeHeight|implemented
299|SE_WakeTheDead|implemented, @Pets, summon one temporary pet from nearby corpses that last a set duration, base: none, limit: none, max: duration (seconds). Note: max range of corpse is 250.
300|SE_Doppelganger|implemented
301|SE_ArcheryDamageModifier|implemented[AA] - increase archery damage by percent
302|SE_FcDamagePctCrit|implemented, @Fc, On Caster, spell damage mod pct, base: min pct, limit: max pct, Note: applied after critical hits has been calculated.
303|SE_FcDamageAmtCrit|implemented, @Fc, On Caster, spell damage mod flat amt, base: amt
304|SE_OffhandRiposteFail|implemented as bonus - enemy cannot riposte offhand attacks
305|SE_MitigateDamageShield|implemented - off hand attacks only (Shielding Resistance)
306|SE_ArmyOfTheDead|implemented, @Pets, summon multiple temporary pets from nearby corpses that last a set duration, base: amount of corpses that a pet can summon from, limit: none, max: duration (seconds). Note: max range of corpse is 250.
307|SE_Appraisal|*not implemented Rogue AA - This ability allows you to estimate the selling price of an item you are holding on your cursor.
308|SE_ZoneSuspendMinion|implemented, @Pet, allow suspended pets to be resummoned upon zoning, base: 1, limit: none, max: none, Calc: Bool
309|SE_GateCastersBindpoint|implemented - Gate to casters bind point
310|SE_ReduceReuseTimer|implemented, @Fc, On Caster, spell and disc reuse time mod by amount, base: milliseconds
311|SE_LimitCombatSkills|implemented, @Ff, Include or exclude combat skills or procs from being focused, base1: 0=Exclude if proc 1=Allow only if proc.
312|SE_Sanctuary|implemented - Places caster at bottom hate list, effect fades if cast cast spell on targets other than self.
313|SE_ForageAdditionalItems|implemented[AA] - chance to forage additional items
314|SE_Invisibility2|implemented - fixed duration invisible
315|SE_InvisVsUndead2|implemented - fixed duration ITU
316|SE_ImprovedInvisAnimals|implemented
317|SE_ItemHPRegenCapIncrease|implemented[AA] - increases amount of health regen gained via items
318|SE_ItemManaRegenCapIncrease|implemented - increases amount of mana regen you can gain via items
319|SE_CriticalHealOverTime|implemented
320|SE_ShieldBlock|implemented - Block attacks with shield
321|SE_ReduceHate|implemented
322|SE_GateToHomeCity|implemented
323|SE_DefensiveProc|implemented
324|SE_HPToMana|implemented
325|SE_NoBreakAESneak|implemented[AA] - [AA Nerves of Steel] increasing chance to remain hidden when they are an indirect target of an AoE spell.
326|SE_SpellSlotIncrease|*not implemented as bonus - increases your spell slot availability
327|SE_MysticalAttune|implemented - increases amount of buffs that a player can have
328|SE_DelayDeath|implemented - increases how far you can fall below 0 hp before you die
329|SE_ManaAbsorbPercentDamage|implemented
330|SE_CriticalDamageMob|implemented
331|SE_Salvage|implemented - chance to recover items that would be destroyed in failed tradeskill combine
332|SE_SummonToCorpse|*not implemented AA - Call of the Wild (Druid/Shaman Res spell with no exp) TOOD: implement this.
333|SE_CastOnRuneFadeEffect|implemented
334|SE_BardAEDot|implemented
335|SE_BlockNextSpellFocus|implemented, @Fc, On Caster, chance to block next spell, base: chance
336|SE_IllusionaryTarget|not used
337|SE_PercentXPIncrease|implemented
338|SE_SummonAndResAllCorpses|implemented
339|SE_TriggerOnCast|implemented, @Fc, On Caster, cast on spell use, base: chance pct limit: spellid
340|SE_SpellTrigger|implemented - chance to trigger spell [Share rolls with 469] All base2 spells share roll chance, only 1 cast.
341|SE_ItemAttackCapIncrease|implemented[AA] - increases the maximum amount of attack you can gain from items.
342|SE_ImmuneFleeing|implemented - stop mob from fleeing
343|SE_InterruptCasting|implemented - % chance to interrupt spells being cast every tic. Cacophony (8272)
344|SE_ChannelChanceItems|implemented[AA] - chance to not have ITEM effects interrupted when you take damage.
345|SE_AssassinateLevel|implemented as bonus - AA Assisination max level to kill
346|SE_HeadShotLevel|implemented[AA] - HeadShot max level to kill
347|SE_DoubleRangedAttack|implemented - chance at an additional archery attack (consumes arrow)
348|SE_LimitManaMin|implemented, @Ff, Mininum mana of spell that can be focused, base1: mana amt
349|SE_ShieldEquipDmgMod|implemented[AA] Increase melee base damage (indirectly increasing hate) when wearing a shield.
350|SE_ManaBurn|implemented - Drains mana for damage/heal at a defined ratio up to a defined maximum amount of mana.
351|SE_PersistentEffect|*not implemented. creates a trap/totem that casts a spell (spell id + base1?) when anything comes near it. can probably make a beacon for this
352|SE_IncreaseTrapCount|*not implemented - looks to be some type of invulnerability? Test ITC (8755)
353|SE_AdditionalAura|*not implemented - allows use of more than 1 aura, aa effect
354|SE_DeactivateAllTraps|*not implemented - looks to be some type of invulnerability? Test DAT (8757)
355|SE_LearnTrap|*not implemented - looks to be some type of invulnerability? Test LT (8758)
356|SE_ChangeTriggerType|not used
357|SE_FcMute|implemented, @Fc, On Caster, prevents spell casting, base: chance pct
358|SE_CurrentManaOnce|implemented
359|SE_PassiveSenseTrap|*not implemented - Invulnerability (Brell's Blessing)
360|SE_ProcOnKillShot|implemented - a buff that has a base1 % to cast spell base2 when you kill a "challenging foe" base3 min level
361|SE_SpellOnDeath|implemented - casts spell on death of buffed
362|SE_PotionBeltSlots|*not implemented[AA] 'Quick Draw' expands the potion belt by one additional available item slot per rank.
363|SE_BandolierSlots|*not implemented[AA] 'Battle Ready' expands the bandolier by one additional save slot per rank.
364|SE_TripleAttackChance|implemented
365|SE_ProcOnSpellKillShot|implemented - chance to trigger a spell on kill when the kill is caused by a specific spell with this effect in it (10470 Venin)
366|SE_GroupShielding|*not implemented[AA] This gives you /shieldgroup
367|SE_SetBodyType|implemented - set body type of base1 so it can be affected by spells that are limited to that type (Plant, Animal, Undead, etc)
368|SE_FactionMod|*not implemented - increases faction with base1 (faction id, live won't match up w/ ours) by base2
369|SE_CorruptionCounter|implemented
370|SE_ResistCorruption|implemented
371|SE_AttackSpeed4|implemented - stackable slow effect 'Inhibit Melee'
372|SE_ForageSkill|implemented[AA] Will increase the skill cap for those that have the Forage skill and grant the skill and raise the cap to those that do not.
373|SE_CastOnFadeEffectAlways|implemented - Triggers if fades after natural duration OR from rune/numhits fades.
374|SE_ApplyEffect|implemented
375|SE_DotCritDmgIncrease|implemented - Increase damage of DoT critical amount
376|SE_Fling|*not implemented - used in 2 test spells  (12945 | Movement Test Spell 1)
377|SE_CastOnFadeEffectNPC|implemented - Triggers only if fades after natural duration (On live these are usually players spells that effect an NPC).
378|SE_SpellEffectResistChance|implemented - Increase chance to resist specific spell effect (base1=value, base2=spell effect id)
379|SE_ShadowStepDirectional|implemented - handled by client
380|SE_Knockdown|implemented - small knock back(handled by client)
381|SE_KnockTowardCaster|*not implemented (Call of Hither) knocks you back to caster (value) distance units infront
382|SE_NegateSpellEffect|implemented, @Debuff, negates specific spell effect benefits for the duration of the debuff and prevent non-duration spell effect from working, base: see NegateSpellEffecttype Enum, limit: SPA id, max: none
383|SE_SympatheticProc|implemented, @Fc, On Caster, cast on spell use, base: variable proc chance on cast time, limit: spellid
384|SE_Leap|implemented - Leap effect, ie stomping leap
385|SE_LimitSpellGroup|implemented, @Ff, Spell group(s) that a spell focus can require or exclude, base1: spellgroup id, Include: Positive Exclude: Negative
386|SE_CastOnCurer|implemented - Casts a spell on the person curing
387|SE_CastOnCure|implemented - Casts a spell on the cured person
388|SE_SummonCorpseZone|implemented - summons a corpse from any zone(nec AA)
389|SE_FcTimerRefresh|implemented, @Fc, On Caster, reset all recast timers, base: 1, Note: Applied from casted spells only
390|SE_FcTimerLockout|implemented, @Fc, On Caster, set a spell to be on recast timer, base: recast duration milliseconds, Note: Applied from casted spells only
391|SE_LimitManaMax|implemented, @Ff, Mininum mana of spell that can be focused, base1: mana amt
392|SE_FcHealAmt|implemented, @Fc, On Caster, spell healing mod flat amt, base: amt
393|SE_FcHealPctIncoming|implemented, @Fc, On Target, heal received mod pct, base: pct, limit: random max pct
394|SE_FcHealAmtIncoming|implemented, @Fc, On Target, heal received mod flat amt, base: amt
395|SE_FcHealPctCritIncoming|implemented, @Fc, On Target, heal received mod pct, base: pct, limit: random max pct
396|SE_FcHealAmtCrit|implemented, @Fc, On Caster, spell healing mod flat amt, base: amt
397|SE_PetMeleeMitigation|implemented[AA] - additional mitigation to your pets. Adds AC
398|SE_SwarmPetDuration|implemented - Affects the duration of swarm pets
399|SE_FcTwincast|implemented - cast 2 spells for every 1
400|SE_HealGroupFromMana|implemented - Drains mana and heals for each point of mana drained
401|SE_ManaDrainWithDmg|implemented - Deals damage based on the amount of mana drained
402|SE_EndDrainWithDmg|implemented - Deals damage for the amount of endurance drained
403|SE_LimitSpellClass|implemented, @Ff, 'Spell Category' using table field 'spell_class' that a spell focus can require or exclude, base1: category type, Include: Positive Exclude: Negative
404|SE_LimitSpellSubclass|implemented, @Ff, 'Spell Category Subclass' using table field 'spell_subclass' that a spell focus can require or exclude, base1: category type, Include: Positive Exclude: Negative
405|SE_TwoHandBluntBlock|implemented - chance to block attacks when using two hand blunt weapons (similiar to shield block)
406|SE_CastonNumHitFade|implemented - casts a spell when a buff fades due to its numhits being depleted
407|SE_CastonFocusEffect|implemented - casts a spell if focus limits are met (ie triggers when a focus effects is applied)
408|SE_LimitHPPercent|implemented - limited to a certain percent of your hp(ie heals up to 50%)
409|SE_LimitManaPercent|implemented - limited to a certain percent of your mana
410|SE_LimitEndPercent|implemented - limited to a certain percent of your end
411|SE_LimitClass|implemented, @Ff, Class(es) that can use the spell focus, base1: class(es), Note: The class value in dbase is +1 in relation to item class value, set as you would item for multiple classes
412|SE_LimitRace|implemented, @Ff, Race that can use the spell focus, base1: race, Note: not used in any known live spells. Use only single race at a time.
413|SE_FcBaseEffects|implemented, @Fc, On Caster, base spell effectiveness mod pct, base: pct
414|SE_LimitCastingSkill|implemented, @Ff, Spell and singing skills(s) that a spell focus can require or exclude, base1: skill id, Include: Positive Exclude: Negative
415|SE_FFItemClass|implemented, @Ff, Limits focuses to be applied only from item click. base1: item ItemType (-1 to include for all ItemTypes,-1000 to exclude clicks from getting the focus, or exclude specific SubTypes or Slots if set), limit: item SubType (-1 for all SubTypes), max: item Slots (bitmask of valid slots, -1 ALL slots), Note: not used on live. See comments in Mob::CalcFocusEffect for more details.
416|SE_ACv2|implemented - New AC spell effect
417|SE_ManaRegen_v2|implemented - New mana regen effect
418|SE_SkillDamageAmount2|implemented - adds skill damage directly to certain attacks
419|SE_AddMeleeProc|implemented - Adds a proc
420|SE_FcLimitUse|implemented, @Fc, On Caster, numhits mod pct, base: pct, Note: not used in any known live spells
421|SE_FcIncreaseNumHits|implemented, @Fc, On Caster, numhits mod flat amt, base: amt
422|SE_LimitUseMin|implemented, @Ff Minium amount of numhits for a spell to be focused, base: numhit amt
423|SE_LimitUseType|implemented,	@Ff Focus will only affect if has this numhits type, base: numhit type
424|SE_GravityEffect|implemented - Pulls/pushes you toward/away the mob at a set pace
425|SE_Display|*not implemented - Illusion: Flying Dragon(21626)
426|SE_IncreaseExtTargetWindow|*not implmented[AA] - increases the capacity of your extended target window
427|SE_SkillProcAttempt|implemented - chance to proc when using a skill(ie taunt)
428|SE_LimitToSkill|implemented, @Procs, limits what combat skills will effect a skill proc, base: skill value, limit: none, max: none
429|SE_SkillProcSuccess|implemented - chance to proc when tje skill in use successfully fires.
430|SE_PostEffect|*not implemented - Fear of the Dark(27641) - Alters vision
431|SE_PostEffectData|*not implemented - Fear of the Dark(27641) - Alters vision
432|SE_ExpandMaxActiveTrophyBen|not used
433|SE_CriticalDotDecay|implemented - increase critical dot chance, effect decays based on level of spell it effects. (12266)
434|SE_CriticalHealDecay|implemented - increase critical heal chance, effect decays based on level of spell it effects.
435|SE_CriticalRegenDecay|implemented - increase critical heal over time chance, effect decays based on level of spell it effects.
436|SE_BeneficialCountDownHold|not used ( 23491 | ABTest Buff Hold)
437|SE_TeleporttoAnchor|*not implemented - Teleport Guild Hall Anchor(33099)
438|SE_TranslocatetoAnchor|*not implemented - Translocate Primary Anchor (27750)
439|SE_Assassinate|implemented[AA] - Assassinate damage
440|SE_FinishingBlowLvl|implemented[AA] - Sets the level Finishing blow can be triggered on an NPC
441|SE_DistanceRemoval|implemented - Buff is removed from target when target moves X amount of distance away from where initially hit.
442|SE_TriggerOnReqTarget|implemented, @SpellTrigger, triggers a spell when Target Requirement conditions are met (see enum SpellRestriction for IDs), base: spellid, limit: SpellRestriction ID, max: none, Note: Usually cast on a target
443|SE_TriggerOnReqCaster|implemented, @SpellTrigger, triggers a spell when Caster Requirement conditions are met (see enum SpellRestriction for IDs), base: spellid, limit: SpellRestriction ID, max: none, Note: Usually self only
444|SE_ImprovedTaunt|implemented - Locks Aggro On Caster and Decrease other Players Aggro by X% on NPC targets below level Y
445|SE_AddMercSlot|*not implemented[AA] - [Hero's Barracks] Allows you to conscript additional mercs.
446|SE_AStacker|implementet - bufff stacking blocker (26219 | Qirik's Watch)
447|SE_BStacker|implemented
448|SE_CStacker|implemented
449|SE_DStacker|implemented
450|SE_MitigateDotDamage|implemented, @Runes, mitigate incoming dot damage by percentage until rune fades, base: percent mitigation, limit: max dmg absorbed per hit, max: rune amt, Note: If placed on item or AA, will provide stackable percent mitigation.
451|SE_MeleeThresholdGuard|implemented  Partial Melee Rune that only is lowered if melee hits are over X amount of damage
452|SE_SpellThresholdGuard|implemented  Partial Spell Rune that only is lowered if spell hits are over X amount of damage
453|SE_TriggerMeleeThreshold|implemented  Trigger effect on X amount of melee damage taken in a single hit
454|SE_TriggerSpellThreshold|implemented  Trigger effect on X amount of spell damage taken in a single hit
455|SE_AddHatePct|implemented  Modify total hate by %
456|SE_AddHateOverTimePct|implemented  Modify total hate by % over time.
457|SE_ResourceTap|implemented  Coverts a percent of dmg from dmg spells(DD/DoT) to hp/mana/end.
458|SE_FactionModPct|implemented  Modifies faction gains and losses by percent.
459|SE_DamageModifier2|implemented - Modifies melee damage by skill type
460|SE_Ff_Override_NotFocusable|implemented, @Fc, Allow spell to be focused event if flagged with 'not_focusable' in spell table, base: 1
461|SE_ImprovedDamage2|implemented, @Fc, On Caster, spell damage mod pct, base: min pct, limit: max pct
462|SE_FcDamageAmt2|implemented, @Fc, On Caster, spell damage mod flat amt, base: amt
463|SE_Shield_Target|464|SE_PC_Pet_Rampage|implemented - Base1 % chance to do rampage for base2 % of damage each melee round
465|SE_PC_Pet_AE_Rampage|implemented - Base1 % chance to do AE rampage for base2 % of damage each melee round
466|SE_PC_Pet_Flurry_Chance|implemented - Base1 % chance to do flurry from double attack hit.
467|SE_DS_Mitigation_Amount|implemented - Modify incoming damage shield damage by a flat amount
468|SE_DS_Mitigation_Percentage|implemented - Modify incoming damage shield damage by percentage
469|SE_Chance_Best_in_Spell_Grp|implemented - Chance to cast highest scribed spell within a spell group. All base2 spells share roll chance, only 1 cast.
470|SE_Trigger_Best_in_Spell_Grp|implemented - Chance to cast highest scribed spell within a spell group. Each spell has own chance.
471|SE_Double_Melee_Round|implemented, @OffBonus, percent chance to repeat primary weapon round with a percent damage modifier, base: pct chance repeat, limit: pct dmg mod, max: none
472|SE_Buy_AA_Rank|implemented,  @Special, Used in AA abilities that have Enable/Disable toggle. Spell on Disabled Rank has this effect in it, base: 1, limit: none, max: none, Note: This will not just buy an AA
473|SE_Double_Backstab_Front|implemented - Chance to double backstab from front
474|SE_Pet_Crit_Melee_Damage_Pct_Owner|implemenetd - Critical damage mod applied to pets from owner
475|SE_Trigger_Spell_Non_Item|implemented - Trigger spell on cast only if not from item click.
476|SE_Weapon_Stance|implemented, @Misc, Apply a specific spell buffs automatically depending 2Hander, Shield or Dual Wield is equipped, base: spellid, base: 0=2H 1=Shield 2=DW, max: none
477|SE_Hatelist_To_Top_Index|Implemented - Chance to be set to top of rampage list
478|SE_Hatelist_To_Tail_Index|Implemented - Chance to be set to bottom of rampage list
479|SE_Ff_Value_Min|implemented, @Ff, Minimum base value of a spell that can be focused, base: spells to be focused base1 value
480|SE_Ff_Value_Max|implemented, @Ff, Max base value of a spell that can be focused, base: spells to be focused base1 value
481|SE_Fc_Cast_Spell_On_Land|implemented, @Fc, On Target, cast spell if hit by spell, base: chance pct, limit: spellid
482|SE_Skill_Base_Damage_Mod|implemented, @OffBonus, modify base melee damage by percent, base: pct, limit: skill(-1=ALL), max: none
483|SE_Fc_Spell_Damage_Pct_IncomingPC|implemented, @Fc, On Target, spell damage taken mod pct, base: min pct, limit: max pct
484|SE_Fc_Spell_Damage_Amt_IncomingPC|implemented, @Fc, On Target, damage taken flat amt, base: amt
485|SE_Ff_CasterClass|implemented, @Ff, Caster of spell on target with a focus effect that is checked by incoming spells must be specified class(es). base1: class(es), Note: Set multiple classes same as would for items
486|SE_Ff_Same_Caster|implemented, @Ff, Caster of spell on target with a focus effect that is checked by incoming spells, base1: 0=Must be different caster 1=Must be same caster
487|SE_Extend_Tradeskill_Cap|488|SE_Defender_Melee_Force_Pct_PC|489|SE_Worn_Endurance_Regen_Cap|implemented, modify worn regen cap, base: amt, limit: none, max: none
490|SE_Ff_ReuseTimeMin|implemented, @Ff, Minimum recast time of a spell that can be focused, base: recast time
491|SE_Ff_ReuseTimeMax|implemented, @Ff, Max recast time of a spell that can be focused, base: recast time
492|SE_Ff_Endurance_Min|implemented, @Ff, Minimum endurance cost of a spell that can be focused, base: endurance cost
493|SE_Ff_Endurance_Max|implemented, @Ff, Max endurance cost of a spell that can be focused, base: endurance cost
494|SE_Pet_Add_Atk|implemented - Bonus on pet owner which gives their pet increased attack stat
495|SE_Ff_DurationMax|implemented, @Ff, Max duration of spell that can be focused, base: tics
496|SE_Critical_Melee_Damage_Mod_Max|implemented - increase or decrease by percent critical damage (not stackable)
497|SE_Ff_FocusCastProcNoBypass|498|SE_AddExtraAttackPct_1h_Primary|implemented, @OffBonus, gives your double attacks a percent chance to perform an extra attack with 1-handed primary weapon, base: chance, limit: amt attacks, max: none
499|SE_AddExtraAttackPct_1h_Secondary|mplemented, @OffBonus, gives your double attacks a percent chance to perform an extra attack with 1-handed secondary weapon, base: chance, limit: amt attacks, max: none
500|SE_Fc_CastTimeMod2|implemented, @Fc, On Caster, cast time mod pct, base: pct, Note: Can reduce to instant cast
501|SE_Fc_CastTimeAmt|implemented, @Fc, On Caster, cast time mod flat amt, base: milliseconds, Note: Can reduce to instant cast
502|SE_Fearstun|implemented - Stun with a max level limit. Normal stun restrictions don't apply.
503|SE_Melee_Damage_Position_Mod|implemented, @OffBonus, modify melee damage by percent if done from Front or Behind opponent, base: pct, limit: 0=back 1=front, max: none
504|SE_Melee_Damage_Position_Amt|implemented, @OffBonus, modify melee damage by flat amount if done from Front or Behind opponent, base: amt, limit: 0=back 1=front, max: none
505|SE_Damage_Taken_Position_Mod|implemented, @DefBonus, modify melee damage by percent if dmg taken from Front or Behind, base: pct, limit: 0=back 1=front, max: none
506|SE_Damage_Taken_Position_Amt|implemented -@DefBonus, modify melee damage by flat amount if dmg taken from your Front or Behind, base: amt, limit: 0=back 1=front, max: none
507|SE_Fc_Amplify_Mod|implemented, @Fc, On Caster, damage-heal-dot mod pct, base: pct
508|SE_Fc_Amplify_Amt|implemented, @Fc, On Caster, damage-heal-dot mod flat amt, base: amt
509|SE_Health_Transfer|implemented - exchange health for damage or healing on a target. ie Lifeburn/Act of Valor
510|SE_Fc_ResistIncoming|implemented, @Fc, On Target, resist modifier, base: amt
511|SE_Ff_FocusTimerMin|implemented, @Ff, sets a recast time until focus can be used again, base: 1, limit: time ms, Note:  ie. limit to 1 trigger every 1.5 seconds
512|SE_Proc_Timer_Modifier|implemented - limits procs per amount of a time based on timer value, base: 1, limit: time ms, Note:, ie limit to 1 proc every 55 seconds)
513|SE_Mana_Max_Percent|514|SE_Endurance_Max_Percent|515|SE_AC_Avoidance_Max_Percent|implemented - stackable avoidance modifier
516|SE_AC_Mitigation_Max_Percent|implemented - stackable defense modifier
517|SE_Attack_Offense_Max_Percent|518|SE_Attack_Accuracy_Max_Percent|implemented - stackable accurary modifer
519|SE_Luck_Amount|520|SE_Luck_Percent|521|SE_Endurance_Absorb_Pct_Damage|implemented - Reduces % of Damage using Endurance, drains endurance at a ratio (ie. 0.05 Endurance per Hit Point)
522|SE_Instant_Mana_Pct|implemented - Increase/Decrease mana by percent of max mana
523|SE_Instant_Endurance_Pct|implemented - Increase/Decrease mana by percent of max endurance
524|SE_Duration_HP_Pct|implemented - Decrease Current Hit Points by % of Total Hit Points per Tick, up to a MAX per tick
525|SE_Duration_Mana_Pct|implemented - Decrease Current Mana by % of Total Mana per Tick, up to a MAX per tick
526|SE_Duration_Endurance_Pct|implemented - Decrease Current Endurance by % of Total Hit Points per Tick, up to a MAX per tick
