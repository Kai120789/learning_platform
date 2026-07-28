import { createSlice } from "@reduxjs/toolkit";
import type { GroupData, GroupResponse, GroupSchema, GroupUser, ShortUserInfo } from "../types/types";
import { getGroupsByStudentId } from "@/entities/group";
import { getGroupsByTutorId } from "@/entities/group";
import { createGroup } from "@/entities/group";
import { deleteGroup } from "@/entities/group";
import { removeUserFromGroup } from "@/entities/group";
import { updateGroup } from "@/entities/group";
import { addUsersToGroup } from "@/entities/group";

const initialState: GroupSchema = {
    data: null,
    isLoading: false,
    error: undefined
};

function mapShortUser(user: ShortUserInfo): GroupUser {
    return {
        id: user.id,
        name: user.name,
        surname: user.surname,
        patronymic: user.patronymic,
        tgUsername: user.tg_username,
    }
}

function mapGroupResponse(group: GroupResponse): GroupData {
    return {
        id: group.id,
        title: group.title,
        description: group.description,
        subject: group.subject,
        users: group.users?.map(mapShortUser) ?? [],
        tutorId: group.tutor_id,
        tgChatId: group.tg_chat_id,
        tgGroupLink: group.tg_group_link,
    }
}

const groupSlice = createSlice({
    name: 'group',
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder.addCase(getGroupsByStudentId.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(getGroupsByStudentId.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getGroupsByStudentId.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data = action.payload.map(mapGroupResponse)
        })
        builder.addCase(getGroupsByTutorId.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(getGroupsByTutorId.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getGroupsByTutorId.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data = action.payload.map(mapGroupResponse)
        })
        builder.addCase(createGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(createGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(createGroup.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            if (!state.data) {
                state.data = []
            }
            state.data.push(mapGroupResponse(action.payload))
        })
        builder.addCase(deleteGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(deleteGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteGroup.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data = state.data?.filter(g => g.id != action.meta.arg) || null
        })

        builder.addCase(removeUserFromGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(removeUserFromGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(removeUserFromGroup.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            const group = state.data?.find(
                g => g.id === action.meta.arg.groupID
            );

            if (group && group.users) {
                group.users = group.users.filter(
                    u => u.id !== action.meta.arg.userID
                );
            }
        })
        builder.addCase(addUsersToGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(addUsersToGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(addUsersToGroup.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''

            const group = state.data?.find(
                g => g.id === action.meta.arg.groupID
            )
            if (!group) return

            const existingIDs = new Set(group.users?.map((user) => user.id))
            const addedUsers = action.meta.arg.users
                .filter((user) => !existingIDs.has(user.id))
                .map(mapShortUser)

            group.users = [...(group.users ?? []), ...addedUsers]
        })
        builder.addCase(updateGroup.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(updateGroup.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(updateGroup.fulfilled, (state, action) => {
            state.isLoading = false;
            state.error = '';

            const group = state.data?.find(g => g.id === action.meta.arg.groupID);

            if (group) {
                group.title = action.payload.title;
                group.description = action.payload.description;
                group.tutorId = action.payload.tutor_id;
                group.tgGroupLink = action.payload.tg_group_link;
                group.tgChatId = action.payload.tg_chat_id;
            }
        });
    }
});

export const { actions: groupActions, reducer: groupReducer } =
    groupSlice;
