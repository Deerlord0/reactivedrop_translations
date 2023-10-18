package main

type achievementData struct {
	id         int32
	bit        uint8
	as2k4Medal int8
	apiName    string
}

var achievements = [...]achievementData{
	{146, 0, -1, "ASW_KILL_WITHOUT_FRIENDLY_FIRE"},
	{146, 1, 2, "ASW_NO_FRIENDLY_FIRE"},
	{146, 2, 3, "ASW_SHIELDBUG"},
	{146, 3, 4, "ASW_GRENADE_MULTI_KILL"},
	{146, 4, 6, "ASW_ACCURACY"},
	{146, 5, 8, "ASW_NO_DAMAGE_TAKEN"},
	{146, 6, 10, "ASW_EGGS_BEFORE_HATCH"},
	{146, 7, -1, "ASW_GRUB_KILLS"},
	{146, 8, 12, "ASW_MELEE_PARASITE"},
	{146, 9, 14, "ASW_MELEE_KILLS"},
	{146, 10, 16, "ASW_BARREL_KILLS"},
	{146, 11, 28, "ASW_INFESTATION_CURING"},
	{146, 12, -1, "ASW_FAST_WIRE_HACKS"},
	{146, 13, -1, "ASW_FAST_COMPUTER_HACKS"},
	{146, 14, -1, "ASW_GROUP_HEAL"},
	{146, 15, 21, "ASW_GROUP_DAMAGE_AMP"},
	{146, 16, 20, "ASW_FAST_RELOADS_IN_A_ROW"},
	{146, 17, -1, "ASW_FAST_RELOAD"},
	{146, 18, 26, "ASW_ALL_HEALING"},
	{146, 19, 29, "ASW_PROTECT_TECH"},
	{146, 20, -1, "ASW_TECH_SURVIVES"},
	{146, 21, 5, "ASW_STUN_GRENADE"},
	{146, 22, -1, "ASW_WELD_DOOR"},
	{146, 23, 7, "ASW_DODGE_RANGER_SHOT"},
	{146, 24, 11, "ASW_BOOMER_KILL_EARLY"},
	{146, 25, 19, "ASW_FREEZE_GRENADE"},
	{146, 26, -1, "ASW_AMMO_RESUPPLY"},
	{146, 27, -1, "ASW_SENTRY_GUN_KILLS"},
	{146, 28, -1, "ASW_RIFLE_KILLS"},
	{146, 29, -1, "ASW_PRIFLE_KILLS"},
	{146, 30, -1, "ASW_AUTOGUN_KILLS"},
	{146, 31, -1, "ASW_SHOTGUN_KILLS"},
	{147, 0, -1, "ASW_VINDICATOR_KILLS"},
	{147, 1, -1, "ASW_PISTOL_KILLS"},
	{147, 2, -1, "ASW_PDW_KILLS"},
	{147, 3, -1, "ASW_TESLA_GUN_KILLS"},
	{147, 4, -1, "ASW_RAILGUN_KILLS"},
	{147, 5, -1, "ASW_FLAMER_KILLS"},
	{147, 6, -1, "ASW_CHAINSAW_KILLS"},
	{147, 7, -1, "ASW_MINIGUN_KILLS"},
	{147, 8, -1, "ASW_SNIPER_RIFLE_KILLS"},
	{147, 9, -1, "ASW_GRENADE_LAUNCHER_KILLS"},
	{147, 10, -1, "ASW_HORNET_KILLS"},
	{147, 11, -1, "ASW_LASER_MINE_KILLS"},
	{147, 12, -1, "ASW_MINE_KILLS"},
	{147, 13, -1, "ASW_UNLOCK_ALL_WEAPONS"},
	{147, 14, -1, "ASW_EASY_CAMPAIGN"},
	{147, 15, -1, "ASW_NORMAL_CAMPAIGN"},
	{147, 16, -1, "ASW_HARD_CAMPAIGN"},
	{147, 17, -1, "ASW_INSANE_CAMPAIGN"},
	{147, 18, -1, "ASW_KILL_GRIND_1"},
	{147, 19, -1, "ASW_KILL_GRIND_2"},
	{147, 20, -1, "ASW_KILL_GRIND_3"},
	{147, 21, -1, "ASW_KILL_GRIND_4"},
	{147, 22, 44, "ASW_SPEEDRUN_LANDING_BAY"},
	{147, 23, 48, "ASW_SPEEDRUN_DESCENT"},
	{147, 24, 45, "ASW_SPEEDRUN_DEIMA"},
	{147, 25, 46, "ASW_SPEEDRUN_RYDBERG"},
	{147, 26, 47, "ASW_SPEEDRUN_RESIDENTIAL"},
	{147, 27, 49, "ASW_SPEEDRUN_SEWER"},
	{147, 28, 50, "ASW_SPEEDRUN_TIMOR"},
	{147, 29, -1, "ASW_CAMPAIGN_NO_DEATHS"},
	{147, 30, 9, "ASW_MISSION_NO_DEATHS"},
	{147, 31, -1, "ASW_IMBA_CAMPAIGN"},
	{148, 0, -1, "ASW_HARDCORE"},
	{148, 1, -1, "ASW_COMBAT_RIFLE_KILLS"},
	{148, 2, -1, "ASW_DEAGLE_KILLS"},
	{148, 3, -1, "ASW_DEVASTATOR_KILLS"},
	{148, 4, -1, "RD_SPEEDRUN_OCS_STORAGE_FACILITY"},
	{148, 5, -1, "RD_SPEEDRUN_OCS_LANDING_BAY_7"},
	{148, 6, -1, "RD_SPEEDRUN_OCS_USC_MEDUSA"},
	{148, 7, -1, "RD_SPEEDRUN_RES_FOREST_ENTRANCE"},
	{148, 8, -1, "RD_SPEEDRUN_RES_RESEARCH_7"},
	{148, 9, -1, "RD_SPEEDRUN_RES_MINING_CAMP"},
	{148, 10, -1, "RD_SPEEDRUN_RES_MINES"},
	{148, 11, -1, "RD_CAMPAIGN_NO_DEATHS_OCS"},
	{148, 12, -1, "RD_CAMPAIGN_NO_DEATHS_RES"},
	{148, 13, -1, "RD_EASY_CAMPAIGN_OCS"},
	{148, 14, -1, "RD_NORMAL_CAMPAIGN_OCS"},
	{148, 15, -1, "RD_HARD_CAMPAIGN_OCS"},
	{148, 16, -1, "RD_INSANE_CAMPAIGN_OCS"},
	{148, 17, -1, "RD_IMBA_CAMPAIGN_OCS"},
	{148, 18, -1, "RD_EASY_CAMPAIGN_RES"},
	{148, 19, -1, "RD_NORMAL_CAMPAIGN_RES"},
	{148, 20, -1, "RD_HARD_CAMPAIGN_RES"},
	{148, 21, -1, "RD_INSANE_CAMPAIGN_RES"},
	{148, 22, -1, "RD_IMBA_CAMPAIGN_RES"},
	{148, 23, -1, "RD_CAMPAIGN_NO_DEATHS_AREA9800"},
	{148, 24, -1, "RD_SPEEDRUN_AREA9800_LZ"},
	{148, 25, -1, "RD_SPEEDRUN_AREA9800_PP1"},
	{148, 26, -1, "RD_SPEEDRUN_AREA9800_PP2"},
	{148, 27, -1, "RD_SPEEDRUN_AREA9800_WL"},
	{148, 28, -1, "RD_EASY_CAMPAIGN_AREA9800"},
	{148, 29, -1, "RD_NORMAL_CAMPAIGN_AREA9800"},
	{148, 30, -1, "RD_HARD_CAMPAIGN_AREA9800"},
	{148, 31, -1, "RD_INSANE_CAMPAIGN_AREA9800"},
	{389, 0, -1, "RD_IMBA_CAMPAIGN_AREA9800"},
	{389, 1, -1, "RD_CAMPAIGN_NO_DEATHS_TFT"},
	{389, 2, -1, "RD_SPEEDRUN_TFT_DESERT_OUTPOST"},
	{389, 3, -1, "RD_SPEEDRUN_TFT_ABANDONED_MAINTENANCE"},
	{389, 4, -1, "RD_SPEEDRUN_TFT_SPACEPORT"},
	{389, 5, -1, "RD_EASY_CAMPAIGN_TFT"},
	{389, 6, -1, "RD_NORMAL_CAMPAIGN_TFT"},
	{389, 7, -1, "RD_HARD_CAMPAIGN_TFT"},
	{389, 8, -1, "RD_INSANE_CAMPAIGN_TFT"},
	{389, 9, -1, "RD_IMBA_CAMPAIGN_TFT"},
	{389, 19, -1, "RD_CAMPAIGN_NO_DEATHS_TIL"},
	{389, 20, -1, "RD_SPEEDRUN_TIL_MIDNIGHT_PORT"},
	{389, 21, -1, "RD_SPEEDRUN_TIL_ROAD_TO_DAWN"},
	{389, 22, -1, "RD_SPEEDRUN_TIL_ARCTIC_INFILTRATION"},
	{389, 23, -1, "RD_SPEEDRUN_TIL_AREA9800"},
	{389, 24, -1, "RD_SPEEDRUN_TIL_COLD_CATWALKS"},
	{389, 25, -1, "RD_SPEEDRUN_TIL_YANAURUS_MINE"},
	{389, 26, -1, "RD_SPEEDRUN_TIL_FACTORY"},
	{389, 27, -1, "RD_SPEEDRUN_TIL_COM_CENTER"},
	{389, 28, -1, "RD_SPEEDRUN_TIL_SYNTEK_HOSPITAL"},
	{389, 29, -1, "RD_EASY_CAMPAIGN_TIL"},
	{389, 30, -1, "RD_NORMAL_CAMPAIGN_TIL"},
	{389, 31, -1, "RD_HARD_CAMPAIGN_TIL"},
	{1462, 0, -1, "RD_INSANE_CAMPAIGN_TIL"},
	{1462, 1, -1, "RD_IMBA_CAMPAIGN_TIL"},
	{1462, 2, -1, "RD_CAMPAIGN_NO_DEATHS_LAN"},
	{1462, 3, -1, "RD_SPEEDRUN_LAN_BRIDGE"},
	{1462, 4, -1, "RD_SPEEDRUN_LAN_SEWER"},
	{1462, 5, -1, "RD_SPEEDRUN_LAN_MAINTENANCE"},
	{1462, 6, -1, "RD_SPEEDRUN_LAN_VENT"},
	{1462, 7, -1, "RD_SPEEDRUN_LAN_COMPLEX"},
	{1462, 8, -1, "RD_EASY_CAMPAIGN_LAN"},
	{1462, 9, -1, "RD_NORMAL_CAMPAIGN_LAN"},
	{1462, 10, -1, "RD_HARD_CAMPAIGN_LAN"},
	{1462, 11, -1, "RD_INSANE_CAMPAIGN_LAN"},
	{1462, 12, -1, "RD_IMBA_CAMPAIGN_LAN"},
	{1462, 25, -1, "RD_CAMPAIGN_NO_DEATHS_PAR"},
	{1462, 26, -1, "RD_EASY_CAMPAIGN_PAR"},
	{1462, 27, -1, "RD_NORMAL_CAMPAIGN_PAR"},
	{1462, 28, -1, "RD_HARD_CAMPAIGN_PAR"},
	{1462, 29, -1, "RD_INSANE_CAMPAIGN_PAR"},
	{1462, 30, -1, "RD_IMBA_CAMPAIGN_PAR"},
	{1462, 31, -1, "RD_SPEEDRUN_PAR_UNEXPECTED_ENCOUNTER"},
	{1517, 0, -1, "RD_SPEEDRUN_PAR_HOSTILE_PLACES"},
	{1517, 1, -1, "RD_SPEEDRUN_PAR_CLOSE_CONTACT"},
	{1517, 2, -1, "RD_SPEEDRUN_PAR_HIGH_TENSION"},
	{1517, 3, -1, "RD_SPEEDRUN_PAR_CRUCIAL_POINT"},
	{1517, 4, -1, "RD_GAS_GRENADE_KILLS"},
	{1517, 5, -1, "RD_HEAVY_RIFLE_KILLS"},
	{1517, 6, -1, "RD_MEDICAL_SMG_KILLS"},
	{1517, 7, -1, "RD_EASY_CAMPAIGN_NH"},
	{1517, 8, -1, "RD_NORMAL_CAMPAIGN_NH"},
	{1517, 9, -1, "RD_HARD_CAMPAIGN_NH"},
	{1517, 10, -1, "RD_INSANE_CAMPAIGN_NH"},
	{1517, 11, -1, "RD_IMBA_CAMPAIGN_NH"},
	{1517, 12, -1, "RD_SPEEDRUN_NH_LOGISTICS_AREA"},
	{1517, 13, -1, "RD_SPEEDRUN_NH_PLATFORM_XVII"},
	{1517, 14, -1, "RD_SPEEDRUN_NH_GROUNDWORK_LABS"},
	{1517, 15, -1, "RD_CAMPAIGN_NO_DEATHS_NH"},
	{1517, 16, -1, "RD_NH_BONUS_OBJECTIVE"},
	{1517, 17, -1, "RD_EASY_CAMPAIGN_BIO"},
	{1517, 18, -1, "RD_NORMAL_CAMPAIGN_BIO"},
	{1517, 19, -1, "RD_HARD_CAMPAIGN_BIO"},
	{1517, 20, -1, "RD_INSANE_CAMPAIGN_BIO"},
	{1517, 21, -1, "RD_IMBA_CAMPAIGN_BIO"},
	{1517, 22, -1, "RD_SPEEDRUN_BIO_OPERATION_X5"},
	{1517, 23, -1, "RD_SPEEDRUN_BIO_INVISIBLE_THREAT"},
	{1517, 24, -1, "RD_SPEEDRUN_BIO_BIOGEN_LABS"},
	{1517, 25, -1, "RD_CAMPAIGN_NO_DEATHS_BIO"},
	{1517, 26, -1, "RD_SPEEDRUN_BONUS_SPC"},
	{1517, 27, -1, "RD_SPEEDRUN_BONUS_RAPTURE"},
	{1517, 28, -1, "RD_SPEEDRUN_BONUS_BUNKER"},
	{1517, 29, -1, "RD_EASY_CAMPAIGN_ACC"},
	{1517, 30, -1, "RD_NORMAL_CAMPAIGN_ACC"},
	{1517, 31, -1, "RD_HARD_CAMPAIGN_ACC"},
	{2557, 0, -1, "RD_INSANE_CAMPAIGN_ACC"},
	{2557, 1, -1, "RD_IMBA_CAMPAIGN_ACC"},
	{2557, 2, -1, "RD_SPEEDRUN_ACC_INFODEP"},
	{2557, 3, -1, "RD_SPEEDRUN_ACC_POWERHOOD"},
	{2557, 4, -1, "RD_SPEEDRUN_ACC_RESCENTER"},
	{2557, 5, -1, "RD_SPEEDRUN_ACC_CONFACILITY"},
	{2557, 6, -1, "RD_SPEEDRUN_ACC_J5CONNECTOR"},
	{2557, 7, -1, "RD_SPEEDRUN_ACC_LABRUINS"},
	{2557, 8, -1, "RD_SPEEDRUN_ACC_COMPLEX"},
	{2557, 9, -1, "RD_CAMPAIGN_NO_DEATHS_ACC"},
	//{2557, 10, -1, "RD_EASY_CAMPAIGN_ADA"},
	//{2557, 11, -1, "RD_NORMAL_CAMPAIGN_ADA"},
	//{2557, 12, -1, "RD_HARD_CAMPAIGN_ADA"},
	//{2557, 13, -1, "RD_INSANE_CAMPAIGN_ADA"},
	//{2557, 14, -1, "RD_IMBA_CAMPAIGN_ADA"},
	//{2557, 15, -1, "RD_SPEEDRUN_ADA_SECTOR_A9"},
	//{2557, 16, -1, "RD_SPEEDRUN_ADA_NEXUS_SUBNODE"},
	//{2557, 17, -1, "RD_SPEEDRUN_ADA_NEON_CARNAGE"},
	//{2557, 18, -1, "RD_SPEEDRUN_ADA_FUEL_JUNCTION"},
	//{2557, 19, -1, "RD_SPEEDRUN_ADA_DARK_PATH"},
	//{2557, 20, -1, "RD_SPEEDRUN_ADA_FORBIDDEN_OUTPOST"},
	//{2557, 21, -1, "RD_SPEEDRUN_ADA_NEW_BEGINNING"},
	//{2557, 22, -1, "RD_SPEEDRUN_ADA_ANOMALY"},
	//{2557, 23, -1, "RD_CAMPAIGN_NO_DEATHS_ADA"},
	{2557, 24, -1, "RD_DIE_IN_MANY_WAYS"},
	{2557, 25, -1, "RD_MA_SCORE_POINTS"},
	{2557, 26, -1, "RD_MA_REACH_VOLCANO_ALIVE"},
	{2557, 27, -1, "RD_MA_VISIT_EACH_ZONE"},
	{2557, 28, -1, "RD_ACC_MUONGEM_KILL"},
	{2557, 29, -1, "RD_EASY_CAMPAIGN_REDUCTION"},
	{2557, 30, -1, "RD_NORMAL_CAMPAIGN_REDUCTION"},
	{2557, 31, -1, "RD_HARD_CAMPAIGN_REDUCTION"},
	{2727, 0, -1, "RD_INSANE_CAMPAIGN_REDUCTION"},
	{2727, 1, -1, "RD_IMBA_CAMPAIGN_REDUCTION"},
	{2727, 2, -1, "RD_SPEEDRUN_REDUCTION_SILENTSTATION"},
	{2727, 3, -1, "RD_SPEEDRUN_REDUCTION_OPSANDSTORM"},
	{2727, 4, -1, "RD_SPEEDRUN_REDUCTION_FALLENCITY"},
	{2727, 5, -1, "RD_SPEEDRUN_REDUCTION_STOWAWAY"},
	{2727, 6, -1, "RD_SPEEDRUN_REDUCTION_INEVITABLEESCALATION"},
	{2727, 7, -1, "RD_SPEEDRUN_REDUCTION_NUCLEARESCORT"},
	{2727, 8, -1, "RD_CAMPAIGN_NO_DEATHS_REDUCTION"},
}

var officialChallenges = [...]string{
	"asbi",
	"difficulty_tier1",
	"difficulty_tier2",
	"level_one",
	"one_hit",
	"energy_weapons",
	"riflemod_classic",
	"rd_first_person",
	"rd_third_person",
}

var officialMissions = [...]string{
	"dm_desert",
	"dm_deima",
	"dm_residential",
	"dm_testlab",
	"asi-jac1-landingbay_01",
	"asi-jac1-landingbay_02",
	"asi-jac2-deima",
	"asi-jac3-rydberg",
	"asi-jac4-residential",
	"asi-jac6-sewerjunction",
	"asi-jac7-timorstation",
	"rd-ocs1storagefacility",
	"rd-ocs2landingbay7",
	"rd-ocs3uscmedusa",
	"rd-res1forestentrance",
	"rd-res2research7",
	"rd-res3miningcamp",
	"rd-res4mines",
	"rd-area9800lz",
	"rd-area9800pp1",
	"rd-area9800pp2",
	"rd-area9800wl",
	"rd-tft1desertoutpost",
	"rd-tft2abandonedmaintenance",
	"rd-tft3spaceport",
	"rd-til1midnightport",
	"rd-til2roadtodawn",
	"rd-til3arcticinfiltration",
	"rd-til4area9800",
	"rd-til5coldcatwalks",
	"rd-til6yanaurusmine",
	"rd-til7factory",
	"rd-til8comcenter",
	"rd-til9syntekhospital",
	"rd-lan1_bridge",
	"rd-lan2_sewer",
	"rd-lan3_maintenance",
	"rd-lan4_vent",
	"rd-lan5_complex",
	"rd-par1unexpected_encounter",
	"rd-par2hostile_places",
	"rd-par3close_contact",
	"rd-par4high_tension",
	"rd-par5crucial_point",
	"rd-nh01_logisticsarea",
	"rd-nh02_platformxvii",
	"rd-nh03_groundworklabs",
	"rd-bio1operationx5",
	"rd-bio2invisiblethreat",
	"rd-bio3biogenlabs",
	"rd-acc1_infodep",
	"rd-acc2_powerhood",
	"rd-acc3_rescenter",
	"rd-acc4_confacility",
	"rd-acc5_j5connector",
	"rd-acc6_labruins",
	"rd-acc_complex",
	"rd-ht-marine_academy",
	"rd-ada_sector_a9",
	"rd-ada_nexus_subnode",
	"rd-ada_neon_carnage",
	"rd-ada_fuel_junction",
	"rd-ada_dark_path",
	"rd-ada_forbidden_outpost",
	"rd-ada_new_beginning",
	"rd-ada_anomaly",
	"rd-reduction1",
	"rd-reduction2",
	"rd-reduction3",
	"rd-reduction4",
	"rd-reduction5",
	"rd-reduction6",
	"rd-bonus_mission1",
	"rd-bonus_mission2",
	"rd-bonus_mission3",
	"rd-bonus_mission4",
	"rd-bonus_mission5",
	"rd-bonus_mission6",
	"rd-bonus_mission7",
	"rd-bonus10_sewrev",
	"rd-bonus12_rydrev",
	"rd-bonus14_cargrev",
	"rd-bonus15_landrev",
}
